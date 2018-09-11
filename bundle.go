package bun

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// HostType represent different types of the hosts.
type HostType string

// File is a safe way to access bundle files.
type File io.ReadCloser

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
func NewBundle(path string) (Bundle, error) {
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
	infos, err := ioutil.ReadDir(b.Path)
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
	Path string
}

// OpenFile tries to open one of the files of the typeName file type.
// If the file is not found, it tries to open it from a correspondent .gzip archive.
// If the .gzip archive is not found as well then returns an error.
// Caller is responsible for closing the file.
// `
func (fo fileOwner) OpenFile(typeName string) (File, error) {
	fileType := GetFileType(typeName)
	notFound := make([]string, 0)
	for _, localPath := range fileType.Paths {
		filePath := path.Join(fo.Path, localPath)
		file, err := os.Open(filePath)
		if err == nil {
			return file, nil // found
		}
		if !os.IsNotExist(err) {
			return nil, err // error
		}
		// not found
		notFound = append(notFound, filePath)
		filePath += ".gz"
		file, err = os.Open(filePath)
		if err != nil {
			if !os.IsNotExist(err) {
				return nil, err // error
			}
			notFound = append(notFound, filePath)
			continue // not found
		}
		r, err := gzip.NewReader(file)
		if err != nil {
			return nil, err // found
		}
		return struct {
			io.Reader
			io.Closer
		}{io.Reader(r), bulkCloser{r, file}}, nil
	}
	return nil, fmt.Errorf("none of the following files are found:\n%v",
		strings.Join(notFound, "\n"))
}

// ReadJSON reads JSON-encoded data from the bundle file and stores the result in
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
