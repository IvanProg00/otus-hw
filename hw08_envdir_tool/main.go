package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must pass a path to directory")
		os.Exit(1)
	}

	path := os.Args[1]

	info, err := os.Stat(path)
	if err != nil {
		pathError, ok := err.(*os.PathError)
		if !ok {
			panic(err)
		}

		fmt.Println(pathError.Err)
		os.Exit(1)
	}

	if !info.IsDir() {
		fmt.Println("passed path is not a directory")
		os.Exit(1)
	}
}
