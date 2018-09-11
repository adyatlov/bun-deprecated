package bun

import (
	"sync"
)

// ContentType defines type of the content in the bundle file.
type ContentType string

const (
	// JSON represents JSON files.
	JSON ContentType = "JSON"
	// Journal represents Journal files.
	Journal = "journal"
	// Dmesg represents dmesg files.
	Dmesg = "dmesg"
)

// FileType Describes a certain kind of file in the bundle.
type FileType struct {
	Name        string
	ContentType ContentType
	// If HostTypes is not empty, then it means that the file belongs to one of
	// the cluster hosts and is relative to the host's directory,
	// otherwise, it's relative to the bundle's root directory.
	Paths       []string
	Description string
	HostTypes   map[HostType]struct{}
}

var (
	fileTypes   = make(map[string]FileType)
	fileTypesMu sync.RWMutex
)

// RegisterFileType adds the file type to the filetype registry. It panics
// if the file type with the same name is already registered.
func RegisterFileType(f FileType) {
	fileTypesMu.Lock()
	defer fileTypesMu.Unlock()
	if _, dup := fileTypes[f.Name]; dup {
		panic("dcosbundle.RegisterFileType: called twice for driver " + f.Name)
	}
	fileTypes[f.Name] = f
}

// GetFileType returns a file type by its name. It panics if the file type
// is not in the registry.
func GetFileType(typeName string) FileType {
	fileTypesMu.RLock()
	defer fileTypesMu.RUnlock()
	fileType, ok := fileTypes[typeName]
	if !ok {
		panic("dcosbundle.RegisterFileType: No such fileType: " + typeName)
	}
	return fileType
}
