package main

import (
	"os"

	"github.com/koss-null/SlothLS/ostools"
)

func main() {
	fileTreeStorage := ostools.NewFStorage("")
	calledPath, err := os.Executable()
	if err != nil {
		panic("Failed to get the path binary was called from")
	}
	fileTreeStorage.ReadPath(calledPath)

}
