package common

import (
	"context"

	"github.com/ipfs/boxo/files"
)

// File is a file
type file struct {
	files.File
	name string
}

// Name returns the name of the file
func (f *file) Name() string {
	return f.name
}

// NewFile creates a new file
func NewFile(name string, data []byte) File {
	return &file{File: files.NewBytesFile(data), name: name}
}

// AddFile adds a file to IPFS
func (c *client) AddFile(file File) (string, error) {
	cidFile, err := c.api.Unixfs().Add(context.Background(), file)
	if err != nil {
		return "", err
	}
	return cidFile.String(), nil
}
