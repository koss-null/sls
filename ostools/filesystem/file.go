package filesystem

import "time"

// var nilFile = File{
// 	path:      "",
// 	name:      "",
// 	weightBit: -1,
// }

// func NilFile() FSObject {
// 	return &nilFile
// }

type File struct {
	path         string
	name         string
	weightBit    int
	lastRevision time.Time
}

func (f *File) CountWeightBit() chan int {
	return make(chan int)
}

func (f *File) IsFile() bool {
	return true
}

func (f *File) IsFolder() bool {
	return false
}
