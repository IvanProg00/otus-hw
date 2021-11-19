package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	from, to      string
	limit, offset int64

	ErrParamFromRequired = errors.New("param \"from\" is required")
	ErrParamToRequired   = errors.New("param \"to\" is required")
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	if err := validateParams(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := Copy(from, to, offset, limit); err != nil {
		panic(err)
	}
}

func validateParams() error {
	if from == "" {
		return ErrParamFromRequired
	}

	if to == "" {
		return ErrParamToRequired
	}

	return nil
}
