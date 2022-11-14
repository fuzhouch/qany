package query

// Media defines interfaces that where code reads data from. It can be
// local file system, S3 or others. It must be used with Format interface
// to get a complete 'where to read' + 'what to read' semantics.
type Media interface {
	// Close method closes Media object cleanly.
	Close()

	// ReaderConvertable method checks whether this media allows
	// opening file as io.Reader interface. Medias like S3 does not
	// support it.
	ReaderConvertable() bool

	// ReaderConvertable method checks whether this media allows
	// opening file as io.Writer interface. Medias like S3 does not
	// support it.
	WriterConvertable() bool

	// Downloadable method checks whether this media supports
	// downloading file to local disk.
	Downloadable() bool

	// Download method takes an action to download a
	// file to local disk for further processing.
	Download(url string) (string, error)
}
