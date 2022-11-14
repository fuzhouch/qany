package query

import (
	"bufio"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/fuzhouch/query/format/csv"
	"github.com/fuzhouch/query/media/fs"
)

func TestQuery_SmokeTest(t *testing.T) {
	mem := afero.NewMemMapFs()
	media, err := fs.New(fs.WithAferoFs(mem))
	assert.Nil(t, err)
	assert.NotNil(t, media)

	format, err := csv.New("test.csv")
	assert.Nil(t, err)

	// Semantics:
	// SELECT * FROM test_csv_file LIMITS 1;
	queryExpr := Select(AllFields()).Top(1)
	stmt, err := queryExpr.From(mem)
	assert.Nil(t, err)

	scanner := bufio.NewScanner(stmt.Reader())
	for scanner.Next() {
		text := scanner.Text()
		assert.NotEqual(t, "", text)
	}
}
