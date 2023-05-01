package main

import (
	"fmt"
	"os"

	"github.com/koss-null/SlothLS/console/printer"
	"github.com/koss-null/SlothLS/ostools/filesystem"
)

func main() {
	fileTreeStorage := filesystem.NewFStorage("")
	calledPath, err := os.Executable()
	if err != nil {
		panic("Failed to get the path binary was called from")
	}

	_, err = fileTreeStorage.ReadPath(calledPath)
	if err != nil {
		panic(err)
	}

	pr := printer.NewPrinter()
	pr.RemoveLine(1)
	pr.PutLine("Test1")
	pr.PutLine("Test2")
	pr.PrintBuffer()
	var s string
	fmt.Scan(&s)
	pr.MoveUp(1)
	pr.RemoveLine(1)
	pr.PutLine("Test1")
	pr.PrintBuffer()
}
