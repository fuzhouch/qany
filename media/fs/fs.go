package fs

import (
	"errors"

	"github.com/spf13/afero"
)

// FileSystem struct represents file reading semantics from local
// harddisk.
type FileSystem struct {
	fs       afero.Fs
	filePath string
}

type FileSystemOption = func(*FileSystem) error

// New function creates a file system media which access local file
// system to load data.
func New(opts ...FileSystemOption) (*FileSystem, error) {
	newObj := &FileSystem{}
	for _, opt := range opts {
		err := opt(newObj)
		if err != nil {
			return nil, err
		}
	}

	if newObj.fs == nil {
		newObj.fs = afero.NewOsFs()
	}

	// filePath should never be empty.
	if newObj.filePath == "" {
		return nil, errors.New("query.media.fs.FilePath.Empty")
	}

	return newObj, nil
}

// WithAferoFs option allows passing an Fs object, allowing underlying
// actual file read/write switching between in-memory memory and true
// file system.
func WithAferoFs(fs afero.Fs) FileSystemOption {
	return func(f *FileSystem) error {
		f.fs = fs
		return nil
	}
}

// WithFilePath option allows setting a file path to read, which is
// always required when processing files at local file system.
func WithFilePath(filePath string) FileSystemOption {
	return func(c *FileSystem) error {
		c.filePath = filePath
		return nil
	}
}
