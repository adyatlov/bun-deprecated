package bun

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// HostType represent different types of the hosts.
type HostType string

const (
	// Master host type
	Master HostType = "master"
	// Agent host type
	Agent = "agent"
	// PublicAgent host type
	PublicAgent = "public agent"
)

// Host represents a host in a DC/OS cluster.
type Host struct {
	Type HostType
	IP   string
	Path string
	fs   Filesystem
}

// OpenFile opens a host-related file. Caller is responsible for closing the file.
func (h Host) OpenFile(fileType string) (io.ReadCloser, error) {
	return openFile(h.fs, h.Path, fileType)
}

// Bundle describes DC/OS diagnostics bundle.
type Bundle struct {
	Path         string
	Hosts        map[string]Host // IP to Host map
	Masters      map[string]Host
	Agents       map[string]Host
	PublicAgents map[string]Host
	fs           Filesystem
}

// NewBundle creates new Bundle
func NewBundle(ctx context.Context, path string) (Bundle, error) {
	b := Bundle{
		Hosts:        make(map[string]Host),
		Masters:      make(map[string]Host),
		Agents:       make(map[string]Host),
		PublicAgents: make(map[string]Host),
	}
	var err error
	b.Path, err = filepath.Abs(path)
	if err != nil {
		log.Printf("Error occurred while detecting absolute path: %v", err)
		return b, err
	}
	if fs, ok := FSFromContext(ctx); ok {
		b.fs = fs
	} else {
		b.fs = OSFS{}
	}
	infos, err := b.fs.ReadDir(b.Path)
	if err != nil {
		return b, err
	}
	const restr = `^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))_(agent_public|agent|master)$`
	re := regexp.MustCompile(restr)
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		groups := re.FindStringSubmatch(info.Name())
		if groups == nil {
			continue
		}
		host := Host{}
		host.IP = groups[1]
		host.Path = filepath.Join(b.Path, info.Name())
		host.fs = b.fs
		switch groups[5] {
		case "master":
			host.Type = Master
			b.Masters[host.IP] = host
		case "agent":
			host.Type = Agent
			b.Agents[host.IP] = host
		case "agent_public":
			host.Type = PublicAgent
			b.PublicAgents[host.IP] = host
		default:
			panic("dcosbundle.NewBundle: unknown host type: " + groups[5])
		}
		b.Hosts[host.IP] = host
	}
	return b, nil
}

// OpenFile opens a file in a root directory of the bundle. The caller is
// responsible for closing the file.
func (b Bundle) OpenFile(fileType string) (File, error) {
	return openFile(b.fs, b.Path, fileType)
}

// openFile tries to open one of the files of the typeName file type.
// If the file is not found, it tries to open it from a correspondent .gzip archive.
// If the .gzip archive is not found returns an error.
func openFile(fs Filesystem, basePath string, typeName string) (file File, err error) {
	fileType, err := GetFileType(typeName)
	if err != nil {
		return
	}
	for _, p := range fileType.Paths {
		filePath := path.Join(basePath, p)
		if file, err = fs.Open(filePath); err == nil {
			return
		}
		if fs.IsNotExist(err) {
			var gzfile File
			if gzfile, err = fs.Open(filePath + ".gz"); err != nil {
				if fs.IsNotExist(err) {
					continue
				} else {
					return
				}
			}
			if file, err = gzip.NewReader(gzfile); err != nil {
				return
			}
			file = struct {
				io.Reader
				io.Closer
			}{io.Reader(file), bulkCloser{file, gzfile}}
			return
		}
	}
	return
}

type bulkCloser []io.Closer

func (bc bulkCloser) Close() error {
	ee := make([]error, 0)
	for _, c := range bc {
		if err := c.Close(); err != nil {
			ee = append(ee, err)
		}
	}
	if len(ee) > 0 {
		var b strings.Builder
		for i, e := range ee {
			b.WriteString(strconv.Itoa(i))
			b.WriteString(": ")
			b.WriteString(e.Error())
			b.WriteString("\n")
		}
		return errors.New(b.String())
	}
	return nil
}
