package bun

import (
	"io"
	"io/ioutil"
	"os"
)

type FileSystem interface {
	// ReadDir reads the directory named by dirname and returns
	// a list of directory entries sorted by filename.
	// It's mocking the io/ioutil.ReadDir.
	ReadDir(string) ([]os.FileInfo, error)

	// Open opens the named file for reading. If successful, methods on
	// the returned file can be used for reading.
	// It's partially mocking the os.Open function.
	Open(string) (File, error)

	// Getwd returns a rooted path name corresponding to the
	// current directory. If the current directory can be
	// reached via multiple paths (due to symbolic links),
	// Getwd may return any one of them.
	// It's mocking os.Getwd.
	Getwd() (string, error)
}

type File interface {
	io.ReadCloser
}

// osFS implements FileSystem
type OSFS struct {
}

func (osfs OSFS) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}
func (osfs OSFS) Open(name string) (File, error) {
	file, err := os.Open(name)
	return File(file), err
}
func (osfs OSFS) Getwd() (string, error) {
	return os.Getwd()
}
