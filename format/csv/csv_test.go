package csv

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/fuzhouch/query"
	"github.com/fuzhouch/query/media/fs"
)

func TestCSV_SmokeTest(t *testing.T) {
	mem := afero.NewMemMapFs()
	media, err := fs.New(fs.WithAferoFs(mem))
	assert.Nil(t, err)
	assert.NotNil(t, media)

	format, err := New(
		WithFilePath("address-book.csv"),
		WithColumns(
			query.STRING("address"),
			query.STRING("name"),
			query.INT64("age")))
	assert.Nil(t, err)
	assert.NotNil(t, format)

	// Semantics:
	// SELECT * FROM test_csv_file LIMITS 1;
	/*
		queryExpr := Select(ALL()).Top(1)
		stmt, err := queryExpr.From(mem)
		assert.Nil(t, err)

		scanner := bufio.NewScanner(stmt.Reader())
		for scanner.Next() {
			text := scanner.Text()
			assert.NotEqual(t, "", text)
		}
	*/
}
