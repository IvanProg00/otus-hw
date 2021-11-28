package main

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
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

		val, err := getEnvValue(dir, fInfo)
		if err != nil {
			return nil, err
		}

		res[fInfo.Name()] = val
	}

	return res, nil
}

func getEnvValue(dir string, fInfo fs.FileInfo) (EnvValue, error) {
	if fInfo.Size() == 0 {
		return EnvValue{
			NeedRemove: true,
		}, nil
	}

	f, err := os.Open(path.Join(dir, fInfo.Name()))
	if err != nil {
		return EnvValue{}, err
	}

	buf := make([]byte, 40)
	val := ""

	for {
		n, err := f.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return EnvValue{}, err
		}

		readed := buf[:n]
		until := bytes.IndexByte(readed, '\n')
		readed = bytes.ReplaceAll(readed, []byte{0}, []byte{'\n'})

		if until >= 0 {
			val += string(readed[:until])
			break
		}
		val += string(readed)
	}

	val = strings.TrimRight(val, " \t")

	return EnvValue{
		Value:      val,
		NeedRemove: false,
	}, nil
}
