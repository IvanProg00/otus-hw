package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	res := make(map[string]EnvValue)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fInfo := range files {
		if fInfo.IsDir() {
			continue
		}

		fmt.Println(fInfo.Name())

		f, err := os.Open(path.Join(dir, fInfo.Name()))
		if err != nil {
			return nil, err
		}

		buf := make([]byte, 4)
		n, err := f.Read(buf)
		val := ""

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		readed := string(buf[:n])
		until := strings.IndexByte(readed, '\n')

		if until < 0 {
			val += readed
		} else {
			val += readed[:until]
		}

		res[fInfo.Name()] = EnvValue{
			Value: val,
			// NeedRemove: false,
		}
	}

	return res, nil
}
