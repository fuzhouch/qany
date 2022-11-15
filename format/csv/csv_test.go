package csv

import (
	"errors"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/fuzhouch/query"
	"github.com/fuzhouch/query/media/fs"
)

func TestCSV_CreateCSV(t *testing.T) {
	mem := afero.NewMemMapFs()
	media, err := fs.New(
		fs.WithAferoFs(mem),
		fs.WithFilePath("address-book.csv"))
	assert.Nil(t, err)
	assert.NotNil(t, media)

	format, err := New(
		WithColumns(
			query.STRING("address"),
			query.STRING("name"),
			query.INT64("age")))
	assert.Nil(t, err)
	assert.NotNil(t, format)
	defer format.Close()
}

func TestCSV_DefaultSchemaBuiltin(t *testing.T) {
	format, err := New(
		WithColumns(
			query.STRING("address"),
			query.STRING("name"),
			query.INT64("age")))
	assert.Nil(t, err)
	assert.NotNil(t, format)
	defer format.Close()
	assert.False(t, format.SchemaBuiltin())
}

func badMockOption() CSVOption {
	return func(*CSV) error {
		return errors.New("query.format.csv.ut.badMockOption")
	}
}

func TestCSV_BadOptionRaiseError(t *testing.T) {
	format, err := New(badMockOption())
	assert.NotNil(t, err)
	assert.Nil(t, format)
	assert.Equal(t, "query.format.csv.ut.badMockOption", err.Error())
}
