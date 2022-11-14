package query

// Format defines file format to read. For example, data can be read from
// either CSV or multi-line JSON. File must work with Media to for a
// 'where to read' + 'what to read' semantics.
type Format interface{}
