package simple

import (
	"fmt"
	"os"

	"github.com/koss-null/funcfrog/pkg/ff"
	"github.com/koss-null/funcfrog/pkg/pipies"
	"github.com/pkg/errors"
)

const (
	BlueColor  = "\033[34m"
	GreenColor = "\033[32m"
	ResetColor = "\033[0m"
)

func ListDir(dir string) error {
	file, err := os.Open(dir)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()

	files, err := file.Readdir(-1)
	if err != nil {
		return errors.WithStack(err)
	}

	dirs := ff.MapFilter(files, func(file os.FileInfo) (string, bool) {
		return GreenColor + file.Name() + "/" + ResetColor, file.IsDir()
	}).Sort(pipies.Less[string])
	fls := ff.MapFilter(files, func(file os.FileInfo) (string, bool) {
		return BlueColor + file.Name() + ResetColor, !file.IsDir()
	}).Sort(pipies.Less[string])

	addNewLine := func(s string) string { return s + "\n" }
	fmt.Printf(*(fls.Map(addNewLine).Reduce(pipies.Sum[string])))
	fmt.Printf(*(dirs.Map(addNewLine).Reduce(pipies.Sum[string])))

	return nil
}

func ListCurDir() error {
	const currentDir = "."

	return ListDir(currentDir)
}
