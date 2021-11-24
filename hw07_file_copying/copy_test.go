package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	require := require.New(t)
	fs, err := os.CreateTemp("testdata", "out.*.txt")
	require.NoError(err)
	fs.Close()
	defer os.Remove(fs.Name())

	const from = "testdata/input.txt"

	tests := []struct {
		offset       int64
		limit        int64
		expectedFile string
	}{
		{
			expectedFile: "testdata/out_offset0_limit0.txt",
		},
		{
			limit:        10,
			expectedFile: "testdata/out_offset0_limit10.txt",
		},
		{
			limit:        1000,
			expectedFile: "testdata/out_offset0_limit1000.txt",
		},
		{
			limit:        10000,
			expectedFile: "testdata/out_offset0_limit10000.txt",
		},
		{
			offset:       100,
			limit:        1000,
			expectedFile: "testdata/out_offset100_limit1000.txt",
		},
		{
			offset:       6000,
			limit:        1000,
			expectedFile: "testdata/out_offset6000_limit1000.txt",
		},
	}

	for _, t := range tests {
		require.NoError(Copy(from, fs.Name(), t.offset, t.limit))

		resFs, err := os.Open(fs.Name())
		require.NoError(err)

		expectedFs, err := os.Open(t.expectedFile)
		require.NoError(err)

		resB := make([]byte, 256)
		expectedB := make([]byte, 256)

		for {
			n1, err1 := resFs.Read(resB)
			if err1 != io.EOF {
				require.NoError(err1)
			}
			n2, err2 := expectedFs.Read(expectedB)
			if err2 != io.EOF {
				require.NoError(err2)
			}

			require.ElementsMatch(resB[:n1], expectedB[:n2])

			if err1 == io.EOF && err2 == io.EOF {
				break
			} else if err1 == io.EOF || err2 == io.EOF {
				require.FailNow("Two files must be finished at same time")
			}
		}
	}
}

func TestCopy_onError(t *testing.T) {
	require := require.New(t)
	f1, err := os.CreateTemp("testdata", "file.*.txt")
	require.NoError(err)
	f1.Close()
	defer os.Remove(f1.Name())

	f2, err := os.CreateTemp("testdata", "file.*.txt")
	require.NoError(err)
	f2.Close()
	defer os.Remove(f2.Name())

	tests := []struct {
		from     string
		to       string
		offset   int64
		limit    int64
		expected error
	}{
		{
			from:     "file_not_exists84398",
			to:       f1.Name(),
			expected: ErrUnsupportedFile,
		},
		{
			from:     f1.Name(),
			to:       f1.Name(),
			expected: ErrUnsupportedFile,
		},
		{
			from:     f1.Name(),
			to:       f2.Name(),
			offset:   1000000,
			expected: ErrOffsetExceedsFileSize,
		},
	}

	for _, t := range tests {
		require.ErrorIs(Copy(t.from, t.to, t.offset, t.limit), t.expected)
	}
}
