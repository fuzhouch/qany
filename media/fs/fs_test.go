package fs

import (
	"errors"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestFS_DefaultAssignOSFs(t *testing.T) {
	fs, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, fs.fs)
}

func TestFS_ExplicitAssignMemOs(t *testing.T) {
	mem := afero.NewMemMapFs()
	fs, err := New(WithAferoFs(mem))
	assert.Nil(t, err)
	assert.NotNil(t, fs.fs)
	assert.Equal(t, mem, fs.fs)
}

func WithErrorTrigger() FileSystemOption {
	return func(*FileSystem) error {
		return errors.New("query.media.fs.ut.IntentionalError")
	}
}

func TestFS_ErrorOnOptionPassing(t *testing.T) {
	mem := afero.NewMemMapFs()
	fs, err := New(WithAferoFs(mem), WithErrorTrigger())
	assert.NotNil(t, err)
	assert.Nil(t, fs)
	assert.Equal(t, "query.media.fs.ut.IntentionalError", err.Error())
}
