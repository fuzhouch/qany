package csv

import (
	"github.com/fuzhouch/query"
)

// CSV struct describes CSV file format.
type CSV struct {
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
