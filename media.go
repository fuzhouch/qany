package query

// Media defines interfaces that where code reads data from. It can be
// local file system, S3 or others. It must be used with Format interface
// to get a complete 'where to read' + 'what to read' semantics.
type Media interface{}
