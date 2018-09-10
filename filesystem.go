package bun

import (
	"io"
	"io/ioutil"
	"os"
)

type key int

// Filesystem abstracts filesystem, primarly, for writting unit-tests.
type Filesystem interface {
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

	// IsNotExist returns a boolean indicating whether the error is known to
	// report that a file or directory does not exist.
	// It's mocking os.IsNotExist
	IsNotExist(err error) bool
}

// File abstraction
type File io.ReadCloser

// OSFS implements Filesystem interface using filesystem.
type OSFS struct {
}

// ReadDir implements Filesystem.ReadDir.
func (osfs OSFS) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

// Open implements Filesystem.Open.
func (osfs OSFS) Open(name string) (File, error) {
	file, err := os.Open(name)
	return File(file), err
}

// Getwd implements Filesystem.Getwd.
func (osfs OSFS) Getwd() (string, error) {
	return os.Getwd()
}

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
func (osfs OSFS) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}
