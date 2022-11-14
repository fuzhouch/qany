package query

// Format defines file format to read. For example, data can be read from
// either CSV or multi-line JSON. File must work with Media to for a
// 'where to read' + 'what to read' semantics.
type Format interface {
	// Close method closes Format object (usually internal file
	// object) cleanly.
	Close()

	// SchemaBuiltin method checks whether this format has schema
	// defined as part of file. For example, a typical CSV file has
	// first column as schema. A multi-line JSON does not have fixed
	// schema across every line.
	SchemaBuiltin() bool

	// SchemaTypesIncluded checks whether the schema definition
	// includes type definition. For example, protobuf schema has
	// types built-in, but CSV does not. In this case, CSV schema
	// should always fall back to string.
	SchemaTypesIncluded() bool
}
