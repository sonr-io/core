package common

import (
	"context"

	"github.com/ipfs/boxo/files"
)

// Folder is a folder
type Folder = files.Directory

// NewFolder creates a new folder
func NewFolder(fs ...File) Folder {
	return files.NewMapDirectory(NewFileMap(fs))
}

// AddFolder adds a folder to IPFS
func (c *client) AddFolder(folder Folder) (string, error) {
	cidFile, err := c.api.Unixfs().Add(context.Background(), folder)
	if err != nil {
		return "", err
	}
	return cidFile.String(), nil
}
