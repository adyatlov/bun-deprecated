package bun

import (
	"errors"
	"os"
	"path"
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
func OpenFile(basePath string, typeName string) (*os.File, error) {
	fileType, err := GetFileType(typeName)
	if err != nil {
		return nil, err
	}
	filePath := path.Join(basePath, fileType.Path)
	return os.Open(filePath)
}
