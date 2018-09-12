package bun

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type fileOwner struct {
	Path string
}

type bulkCloser []io.Closer

func (bc bulkCloser) Close() error {
	e := []string{}
	for _, c := range bc {
		if err := c.Close(); err != nil {
			e = append(e, err.Error())
		}
	}
	if len(e) > 0 {
		return errors.New(strings.Join(e, "\n"))
	}
	return nil
}

// OpenFile opens the files of the typeName file type.
// If the file is not found, it tries to open it from a correspondent .gzip archive.
// If the .gzip archive is not found as well then returns an error.
// Caller is responsible for closing the file.
func (fo fileOwner) OpenFile(typeName string) (File, error) {
	fileType := GetFileType(typeName)
	// TODO: check if the FileType is compatible with the owner (bundle or host).
	notFound := []string{}
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
		// found
		r, err := gzip.NewReader(file)
		if err != nil {
			return nil, err // error
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
	// TODO: check if FileType is JSON
	file, err := fo.OpenFile(typeName)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("bun.fileOwner.ReadJSON: Cannot close file: %v", err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (fo fileOwner) FindFirstLine(typeName string, substr string) (l string, n int, err error) {
	file, err := fo.OpenFile(typeName)
	if err != nil {
		return
	}
	return findFirstLine(file, substr)
}

func findFirstLine(r io.Reader, substr string) (l string, n int, err error) {
	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		if strings.Contains(line, substr) {
			l = line
			n = i
			return
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}
	// Not found
	return
}
