package test

import (
	"os"

	"github.com/adyatlov/bun/fs"
)

type FileSystem struct {
}

func (fs *FileSystem) ReadDir(name string) ([]os.FileInfo, error) {
	return nil, nil
}

func (fs *FileSystem) Open(name string) (fs.File, error) {
	return nil, nil
}
