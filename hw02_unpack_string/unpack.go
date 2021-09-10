package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var b strings.Builder
	isPastLetter := false
	var pastLetter rune

	for _, c := range str {
		if unicode.IsDigit(c) {
			if !isPastLetter {
				return "", ErrInvalidString
			}

			numRep, err := strconv.Atoi(string(c))
			if err != nil {
				return "", ErrInvalidString
			}

			b.WriteString(strings.Repeat(string(pastLetter), numRep))
			isPastLetter = false
		} else {
			if isPastLetter {
				b.WriteRune(pastLetter)
			}
			pastLetter = c
			isPastLetter = true
		}
	}

	if isPastLetter {
		b.WriteRune(pastLetter)
	}

	return b.String(), nil
}
