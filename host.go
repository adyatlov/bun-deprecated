package bun

import "io"

type HostType string

const (
	Master      HostType = "master"
	Agent                = "agent"
	PublicAgent          = "public agent"
)

type Host struct {
	Type HostType
	IP   string
	Path string
}

// OpenFile opens a host-related file. Caller is responsible for closing the file.
func (h Host) OpenFile(fileType string) (io.ReadCloser, error) {
	return OpenFile(h.Path, fileType)
}
