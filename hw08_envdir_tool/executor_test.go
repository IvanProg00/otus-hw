package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		cmd          []string
		prevEnv      map[string]string
		env          Environment
		expected     int
		expectedEnvs map[string]string
	}{
		{
			cmd:      []string{"echo", "hello"},
			expected: 0,
		},
		{
			cmd:      []string{"cat", "blablabla"},
			expected: 1,
		},
		{
			cmd: []string{"echo", "hello", "world"},
			env: Environment{
				"hi": EnvValue{Value: "Hello"},
			},
			expected: 0,
			expectedEnvs: map[string]string{
				"hi": "Hello",
			},
		},
		{
			cmd: []string{"cat", "-n", "incorrect_file"},
			env: Environment{
				"one": EnvValue{Value: "1"},
				"two": EnvValue{Value: "	2"},
				"three": EnvValue{Value: "  3"},
			},
			expected: 1,
			expectedEnvs: map[string]string{
				"one": "1",
				"two": "	2",
				"three": "  3",
			},
		},
		{
			cmd: []string{"echo", "-n", "value"},
			env: Environment{
				"one": EnvValue{NeedRemove: true},
				"two": EnvValue{Value: "2"},
			},
			prevEnv: map[string]string{
				"one":  "hello world",
				"four": "4",
			},
			expected: 0,
			expectedEnvs: map[string]string{
				"one":  "",
				"two":  "2",
				"four": "4",
			},
		},
	}

	require := require.New(t)
	for _, t := range tests {
		for key, val := range t.prevEnv {
			require.NoError(os.Setenv(key, val))
		}

		code := RunCmd(t.cmd, t.env)
		require.Equal(t.expected, code)

		for key, val := range t.expectedEnvs {
			envVal := os.Getenv(key)
			require.Equal(val, envVal)
		}
	}
}
