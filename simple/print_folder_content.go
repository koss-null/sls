package simple

import (
	"fmt"
	"os"
	"strings"

	"github.com/koss-null/funcfrog/pkg/ff"
	"github.com/koss-null/funcfrog/pkg/pipe"
	"github.com/koss-null/funcfrog/pkg/pipies"
	"github.com/pkg/errors"
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

	dirs := pipe.Map(
		ff.Filter(files, func(file *os.FileInfo) bool { return (*file).IsDir() }),
		os.FileInfo.Name,
	).
		Map(addFolderIcon).
		Map(addFolderColor).
		Sort(pipies.Less[string])

	fls := pipe.Map(
		ff.Filter(files, func(file *os.FileInfo) bool { return !(*file).IsDir() }),
		os.FileInfo.Name,
	).
		Map(addFileIcon).
		Map(addFileColor).
		Sort(pipies.Less[string])

	addTabs := func(s string) string { return s + "\t" }
	fmt.Printf(*(fls.Map(addTabs).Reduce(pipies.Sum[string])))
	fmt.Println()
	fmt.Printf(*(dirs.Map(addTabs).Reduce(pipies.Sum[string])))
	fmt.Println()

	return nil
}

func ListCurDir() error {
	const currentDir = "."

	return ListDir(currentDir)
}

func addFolderIcon(s string) string {
	const icon = "📁"
	const separator = ""
	return icon + separator + s
}

func addFolderColor(s string) string {
	const greenColor = "\033[32m"
	return addColor(s, greenColor)
}

func addFileIcon(s string) string {
	const separator = ""
	const extensionSeparator = "."
	splittedFileName := strings.Split(s, extensionSeparator)
	return iconByExtension(splittedFileName[len(splittedFileName)-1]) + separator + s
}

func iconByExtension(ext string) string {
	const (
		fileIcon   = "📄"
		goIcon     = "🐹"
		pythonIcon = "🐍"
		cIcon      = "<C>"
		jsIcon     = "<JS>"
		shIcon     = "🐚"
		configIcon = `⚙️ `
	)
	switch ext {
	case "go", "mod":
		return goIcon
	case "py", "wheel", "__init__":
		return pythonIcon
	case "c", "cpp", "h":
		return cIcon
	case "js":
		return jsIcon
	case "cfg", "yaml", "toml", "conf", "json":
		return configIcon
	case "sh", "bash":
		return shIcon
	default:
		return fileIcon
	}
}

func addFileColor(s string) string {
	const blueColor = "\033[34m"
	return addColor(s, blueColor)
}

func addColor(s, color string) string {
	const ResetColor = "\033[0m"
	return color + s + ResetColor
}
