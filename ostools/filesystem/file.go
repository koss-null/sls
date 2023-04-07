package filesystem

import "time"

var nilFile = File{
	path:      "",
	name:      "",
	weightBit: -1,
}

func NilFile() FSObject {
	return &nilFile
}

type File struct {
	path         string
	name         string
	weightBit    int
	lastRevision time.Time
}
