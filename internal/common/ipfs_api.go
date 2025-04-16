package common

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	iface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// IPFS represents a wrapper interface abstracting the localhost api
type IPFS interface {
	// Add adds a file to IPFS
	Add(data []byte) (string, error)

	// AddFile adds a file to IPFS
	AddFile(file File) (string, error)

	// AddFolder adds a folder to IPFS
	AddFolder(folder Folder) (string, error)

	// Exists returns true if the file exists
	Exists(cid string) (bool, error)

	// Get returns the file contents
	Get(cid string) ([]byte, error)

	// IsPinned returns true if the file is pinned
	IsPinned(ipns string) (bool, error)

	// Ls returns a list of files in the folder
	Ls(cid string) ([]string, error)

	// Pin pins a file
	Pin(cid string, name string) error

	// Unpin unpins a file
	Unpin(cid string) error
}

// File is a file
type File interface {
	files.File
	Name() string
}

// NewFileMap creates a map of files
func NewFileMap(vs []File) map[string]files.Node {
	m := make(map[string]files.Node)
	for _, f := range vs {
		m[f.Name()] = f
	}
	return m
}

type client struct {
	api *rpc.HttpApi
}

// NewIPFS creates a new IPFS client
func NewIPFS() (IPFS, error) {
	api, err := rpc.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return &client{api: api}, nil
}

// Add adds a file to IPFS
func (c *client) Add(data []byte) (string, error) {
	file := files.NewBytesFile(data)
	cidFile, err := c.api.Unixfs().Add(context.Background(), file)
	if err != nil {
		return "", err
	}
	return cidFile.String(), nil
}

// Get returns the file contents
func (c *client) Get(cid string) ([]byte, error) {
	p, err := path.NewPath(cid)
	if err != nil {
		return nil, err
	}
	node, err := c.api.Unixfs().Get(context.Background(), p)
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("unexpected node type: %T", node)
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// IsPinned returns true if the file is pinned
func (c *client) IsPinned(ipns string) (bool, error) {
	_, err := c.api.Name().Resolve(context.Background(), ipns)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Exists returns true if the file exists
func (c *client) Exists(cid string) (bool, error) {
	p, err := path.NewPath(cid)
	if err != nil {
		return false, err
	}
	_, err = c.api.Block().Stat(context.Background(), p)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Pin pins a file
func (c *client) Pin(cid string, name string) error {
	p, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	return c.api.Pin().Add(context.Background(), p, options.Pin.Name(name))
}

// Unpin unpins a file
func (c *client) Unpin(cid string) error {
	p, err := path.NewPath(cid)
	if err != nil {
		return err
	}
	return c.api.Pin().Rm(context.Background(), p)
}

// Ls returns a list of files in the folder
func (c *client) Ls(cid string) ([]string, error) {
	p, err := path.NewPath(cid)
	if err != nil {
		return nil, err
	}
	dirChan := make(chan iface.DirEntry)
	files := make([]string, 0)
	lsErr := make(chan error, 1)
	go func() {
		lsErr <- c.api.Unixfs().Ls(context.Background(), p, dirChan)
	}()
	for dirEnt := range dirChan {
		files = append(files, dirEnt.Name)
	}
	if err := <-lsErr; err != nil {
		return nil, err
	}
	return files, nil
}
