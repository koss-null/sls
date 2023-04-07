package filesystem

import "sync"

type globalFSStorage struct {
	path2Object map[string]FSObject
}

var (
	gfs globalFSStorage
	mx  sync.Mutex
)

func ReadPath(path string) (FSObject, error) {
	mx.Lock()
	defer mx.Unlock()

	obj, found := gfs.path2Object[path]
	if found && !needUpgrade(obj) {
		return obj, nil
	}

	obj, err := readPath(path)
	if err != nil {
		return NilFile(), err
	}
	gfs.path2Object[path] = obj
}

func readPath(path string) (FSObject, error) {

}

func needUpgrade(obj FSObject) bool {
	// TODO: impl
	return false
}
