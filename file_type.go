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

// FileType Describes a kind of files in the bundle (e.g. dcos-marathon.service).
// The empty HostTypes set means that files of this type belong to one of
// the cluster hosts and the Paths are relative to the host directory.
// Otherwise, the Paths are relative to the root directory of the bundle.
type FileType struct {
	Name        string
	ContentType ContentType
	Paths       []string
	Description string
	// HostTypes defines on which host types this file can be found.
	// For example, dcos-marathon.service file can be found only on the masters.
	HostTypes []HostType
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
	hostTypes := make(map[HostType]struct{})
	for _, t := range f.HostTypes {
		if _, ok := hostTypes[t]; ok {
			panic("dcosbundle.RegisterFileType: duplicate HostType: " + t)
		}
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
