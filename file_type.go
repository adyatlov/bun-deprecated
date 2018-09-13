package bun

import (
	"fmt"
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
type FileType struct {
	Name        string
	ContentType ContentType
	Paths       []string
	Description string
	// DirTypes defines on which host types this file can be found.
	// For example, dcos-marathon.service file can be found only on the masters.
	DirTypes []DirType
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
		panic(fmt.Sprintf("bun.RegisterFileType: called twice for file type %v", f.Name))
	}
	dirTypes := make(map[DirType]struct{})
	for _, t := range f.DirTypes {
		if _, ok := dirTypes[t]; ok {
			panic(fmt.Sprintf("bun.RegisterFileType: duplicate DirType: %v in file type %v", t, f.Name))
		}
		dirTypes[t] = struct{}{}
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
		panic(fmt.Sprintf("bun.RegisterFileType: No such fileType: %v", typeName))
	}
	return fileType
}
