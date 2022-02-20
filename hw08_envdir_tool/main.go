package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must pass path and command")
		os.Exit(1)
	} else if len(os.Args) < 3 {
		fmt.Println("you must pass a command")
		os.Exit(1)
	}

	path := os.Args[1]
	if path == "" {
		fmt.Println("path can not be empty")
	}

	env, err := ReadDir(path)
	if err != nil {
		var errorPath *fs.PathError

		ok := errors.As(err, &errorPath)
		if !ok {
			fmt.Println(errorPath.Err)
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	code := RunCmd(os.Args[2:], env)
	os.Exit(code)
}
