package main

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == toPath {
		return ErrUnsupportedFile
	}

	fInfo, err := os.Stat(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	if fInfo.Size() < offset {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 {
		limit = fInfo.Size()
	}
	if offset+limit > fInfo.Size() {
		limit = fInfo.Size() - offset
	}

	srcF, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer srcF.Close()
	srcF.Seek(offset, io.SeekStart)

	dstF, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer dstF.Close()

	reader := io.LimitReader(srcF, limit)
	writer := bufio.NewWriter(dstF)

	bar := pb.Full.Start64(limit)
	defer bar.Finish()
	barReader := bar.NewProxyReader(reader)

	_, err = io.Copy(writer, barReader)
	if err != nil {
		return ErrUnsupportedFile
	}

	return nil
}
