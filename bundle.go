package bun

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	fileOwner
}

// Bundle describes DC/OS diagnostics bundle.
type Bundle struct {
	Hosts        map[string]Host // IP to Host map
	Masters      map[string]Host
	Agents       map[string]Host
	PublicAgents map[string]Host
	fileOwner
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

type fileOwner struct {
	fs   Filesystem
	Path string
}

// OpenFile tries to open one of the files of the typeName file type.
// If the file is not found, it tries to open it from a correspondent .gzip archive.
// If the .gzip archive is not found as well then returns an error.
// Caller is responsible for closing the file.
func (fo fileOwner) OpenFile(typeName string) (file File, err error) {
	var fileType FileType
	fileType, err = GetFileType(typeName)
	if err != nil {
		return
	}
	notFound := make([]string, 0)
	for _, localPath := range fileType.Paths {
		filePath := path.Join(fo.Path, localPath)
		if file, err = fo.fs.Open(filePath); err == nil {
			return // found
		}
		if fo.fs.IsNotExist(err) { // not found
			notFound = append(notFound, filePath)
			var gzfile File
			filePath += ".gz"
			if gzfile, err = fo.fs.Open(filePath); err != nil {
				if fo.fs.IsNotExist(err) {
					notFound = append(notFound, filePath)
					continue // not found
				} else {
					return // error
				}
			}
			if file, err = gzip.NewReader(gzfile); err != nil {
				return //error
			}
			file = struct {
				io.Reader
				io.Closer
			}{io.Reader(file), bulkCloser{file, gzfile}}
			return
		}
	}
	err = fmt.Errorf("none of the following files are found:\n%v",
		strings.Join(notFound, "\n"))
	return
}

// ReadJSON parses the JSON-encoded data from the file and stores the result in
// the value pointed to by v.
func (fo fileOwner) ReadJSON(typeName string, v interface{}) error {
	// TODO: check if fileType is JSON
	file, err := fo.OpenFile(typeName)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error when closing file health: %v", err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
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
