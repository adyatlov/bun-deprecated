package bun

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

type ContentType string

const (
	JSON    ContentType = "JSON"
	Journal             = "journal"
	Dmesg               = "dmesg"
)

type FileType struct {
	Name        string
	ContentType ContentType
	// If HostTypes is not empty, then Path is relative to the host's directory,
	// otherwise, it's relative to the bundle's root directory.
	Path        string
	Description string
	HostTypes   map[HostType]struct{}
}

var (
	fileTypes   = make(map[string]FileType)
	fileTypesMu sync.RWMutex
)

func RegisterFileType(f FileType) {
	fileTypesMu.Lock()
	defer fileTypesMu.Unlock()
	if _, dup := fileTypes[f.Name]; dup {
		panic("dcosbundle.RegisterFileType: called twice for driver " + f.Name)
	}
	fileTypes[f.Name] = f
}

func GetFileType(typeName string) (FileType, error) {
	fileTypesMu.RLock()
	defer fileTypesMu.RUnlock()
	fileType, ok := fileTypes[typeName]
	if !ok {
		return fileType, errors.New("No such fileType: " + typeName)
	}
	return fileType, nil
}

// OpenFile opens bundle file.  Caller is responsible for closing the file.
func OpenFile(basePath string, typeName string) (file File, err error) {
	fileType, err := GetFileType(typeName)
	if err != nil {
		return
	}
	filePath := path.Join(basePath, fileType.Path)
	file, err = os.Open(filePath)
	if os.IsNotExist(err) {
		var gzfile File
		if gzfile, err = os.Open(filePath + ".gz"); err != nil {
			return
		}
		if file, err = gzip.NewReader(gzfile); err != nil {
			return
		}
		file = struct {
			io.Reader
			io.Closer
		}{io.Reader(file),
			bulkCloser{file, gzfile}}
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
