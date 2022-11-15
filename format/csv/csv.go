package csv

import (
	"errors"

	"github.com/fuzhouch/query"
)

// CSV struct describes CSV file format.
type CSV struct {
	filePath        string
	useSchemaInFile bool
	columns         []query.Column
}

type CSVOption = func(*CSV) error

// New function creates a new CSV file format descriptor.
func New(opts ...CSVOption) (*CSV, error) {
	newObj := &CSV{}
	for _, opt := range opts {
		err := opt(newObj)
		if err != nil {
			return nil, err
		}
	}

	// filePath should never be empty.
	if newObj.filePath == "" {
		return nil, errors.New("query.media.csv.FilePath.Empty")
	}
	return newObj, nil
}

// WithColumns option specifies column settings from code instead of
// reading from header. It's useful when parsing intermediate files from
// other streams.
func WithColumns(columns ...query.Column) CSVOption {
	return func(c *CSV) error {
		c.columns = make([]query.Column, len(columns), len(columns))
		for i, col := range columns {
			c.columns[i] = col
		}
		return nil
	}
}

func WithFilePath(filePath string) CSVOption {
	return func(c *CSV) error {
		c.filePath = filePath
		return nil
	}
}

// -------------------------------------------------------------------
// Methods implementing Format interface
// -------------------------------------------------------------------

func (c *CSV) Close() {
}

// SchemaBuiltin method checks whether CSV file uses schema defined
// by first line. Can be false if input is wrong.
func (c *CSV) SchemaBuiltin() bool {
	return c.useSchemaInFile
}
