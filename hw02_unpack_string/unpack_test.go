package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "Прив4ет", expected: "Приввввет"},
		{input: " 2", expected: "  "},
		{input: "😀4", expected: "😀😀😀😀"},
		{input: "\n8", expected: "\n\n\n\n\n\n\n\n"},
		{input: "bla5\n2hi4k", expected: "blaaaaa\n\nhiiiik"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `\\\\\\\5`, expected: `\\\5`},
		{input: `\4\5\6\2\4`, expected: `45624`},
		{input: `\4\53`, expected: `4555`},
		{input: "\n\\4", expected: "\n4"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc", "45", "aaa10b", "abcde88", "8", "8😀", "fk4\n4f89f4", `4\5\4\6`, `\4\565`,
		`qw\ne`, "\n\\b",
	}

	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
