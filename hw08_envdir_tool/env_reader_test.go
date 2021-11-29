package main

import (
	"io/fs"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		dir      string
		expected Environment
	}{
		{
			dir:      "empty",
			expected: Environment{},
		},
		{
			dir: "simple",
			expected: Environment{
				"one": EnvValue{Value: "1"},
				"two": EnvValue{Value: "    2"},
				"three": EnvValue{Value: "	3"},
				"four": EnvValue{Value: "4"},
				"five": EnvValue{Value: "  5"},
				"six": EnvValue{Value: "	6"},
			},
		},
		{
			dir: "delete",
			expected: Environment{
				"delete1":     EnvValue{NeedRemove: true},
				"delete2":     EnvValue{NeedRemove: true},
				"not_delete1": EnvValue{},
				"not_delete2": EnvValue{},
			},
		},
		{
			dir: "specific",
			expected: Environment{
				"BAR":  EnvValue{Value: "WORLD\n\n\nLALA\nend"},
				"FOO":  EnvValue{Value: "hello"},
				"TEXT": EnvValue{Value: ""},
			},
		},
	}

	for _, t := range tests {
		res, err := ReadDir(path.Join("testdata", "env_reader", t.dir))
		require.NoError(err)
		require.Equal(t.expected, res)
	}
}

func TestReadDir_error(t *testing.T) {
	require := require.New(t)

	tests := []struct {
		dir      string
		expected error
	}{
		{
			dir:      "not_exists_directory_4839",
			expected: fs.ErrNotExist,
		},
	}

	for _, t := range tests {
		_, err := ReadDir(t.dir)
		require.ErrorIs(err, t.expected)
	}
}
