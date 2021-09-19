package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type unpackStruct struct {
	b            strings.Builder
	isPastLetter bool
	isPastSlash  bool
	pastLetter   rune
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	unpackS := unpackStruct{
		isPastLetter: false,
		isPastSlash:  false,
	}

	for _, c := range str {
		if unicode.IsDigit(c) && !unpackS.isPastSlash {
			if !unpackS.isPastLetter {
				return "", ErrInvalidString
			}

			numRep, _ := strconv.Atoi(string(c))
			unpackS.b.WriteString(strings.Repeat(string(unpackS.pastLetter), numRep))
			unpackS.isPastLetter = false
		} else if err := ifLetter(&unpackS, c); err != nil {
			return "", ErrInvalidString
		}
	}

	if unpackS.isPastLetter {
		unpackS.b.WriteRune(unpackS.pastLetter)
	}

	return unpackS.b.String(), nil
}

func ifLetter(unpackS *unpackStruct, currentRune rune) error {
	if unpackS.isPastSlash && !unicode.IsDigit(currentRune) && currentRune != '\\' {
		return ErrInvalidString
	}

	if unpackS.isPastLetter && !unpackS.isPastSlash {
		unpackS.b.WriteRune(unpackS.pastLetter)
	}

	if unpackS.isPastSlash {
		unpackS.isPastSlash = false
	} else if currentRune == '\\' {
		unpackS.isPastSlash = true
	}

	unpackS.pastLetter = currentRune
	unpackS.isPastLetter = true
	return nil
}
