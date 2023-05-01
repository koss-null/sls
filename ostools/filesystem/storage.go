package filesystem

import "sync"

type FStorage struct {
	path2Object map[string]FSObject
	mx          *sync.Mutex
}

func NewFStorage(cachePaths ...string) FStorage {
	if len(cachePaths) > 0 {
		totalCachedData := make(map[string]FSObject)
		for _, path := range cachePaths {
			cachedData, err := readCache(path)
			if err != nil {
				// TODO write the error into a log file
				cachedData = make(map[string]FSObject)
			}
			totalCachedData = joinMap(totalCachedData, cachedData)
		}

		return FStorage{
			totalCachedData,
			&sync.Mutex{},
		}
	}

	return FStorage{
		make(map[string]FSObject),
		&sync.Mutex{},
	}
}

func (s FStorage) ReadPath(path string) (*FSObject, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	obj, found := s.path2Object[path]
	if found && !needUpgrade(obj) {
		return &obj, nil
	}

	objs, err := readPath(path)
	if err != nil {
		return nil, err
	}
	for i := range objs {
		s.path2Object[objs[i].Path()] = objs[i]
	}
	return &obj, nil
}

func readPath(path string) ([]FSObject, error) {
	return nil, nil
}

func readCache(path string) (map[string]FSObject, error) {
	// TODO: implement
	return make(map[string]FSObject), nil
}

func needUpgrade(obj FSObject) bool {
	// TODO: impl
	return false
}

func joinMap(mx map[string]FSObject, my map[string]FSObject) map[string]FSObject {
	if len(mx) < len(my) {
		mx, my = my, mx
	}

	for k, v := range my {
		if _, ok := mx[k]; !ok {
			mx[k] = v
		}
	}

	return mx
}
