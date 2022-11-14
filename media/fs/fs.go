package fs

import "github.com/spf13/afero"

// FileSystem struct represents file reading semantics from local
// harddisk.
type FileSystem struct {
	fs afero.Fs
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
