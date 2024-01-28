package main

import (
	"fmt"
	"os"

	"github.com/koss-null/sls/simple"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println(len(args))
		err := simple.ListDir(args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		return
	}

	err := simple.ListCurDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
