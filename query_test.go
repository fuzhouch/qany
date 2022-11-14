package query

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/fuzhouch/query/media/fs"
)

func TestQuery_SmokeTest(t *testing.T) {
	mem := afero.NewMemMapFs()
	fs, err := fs.New(fs.WithAferoFs(mem))
	assert.Nil(t, err)
	assert.NotNil(t, fs)
}
