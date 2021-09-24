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
	isSlash      bool
	pastLetter   rune
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	unpackS := unpackStruct{
		isPastLetter: false,
		isSlash:      false,
	}

	for _, c := range str {
		if unicode.IsDigit(c) && !unpackS.isSlash {
			if !unpackS.isPastLetter {
				return "", ErrInvalidString
			}

			numRep, err := strconv.Atoi(string(c))
			if err != nil {
				return "", ErrInvalidString
			}

			unpackS.b.WriteString(strings.Repeat(string(unpackS.pastLetter), numRep))
			unpackS.isPastLetter = false
		} else if err := validateIsLetterOrSlash(&unpackS, c); err != nil {
			return "", err
		}
	}

	if unpackS.isPastLetter {
		unpackS.b.WriteRune(unpackS.pastLetter)
	}

	return unpackS.b.String(), nil
}

func validateIsLetterOrSlash(unpackS *unpackStruct, currentRune rune) error {
	if unpackS.isSlash {
		if !unicode.IsDigit(currentRune) && currentRune != '\\' {
			return ErrInvalidString
		}
		unpackS.isSlash = false
	} else {
		if unpackS.isPastLetter {
			unpackS.b.WriteRune(unpackS.pastLetter)
		}
		if currentRune == '\\' {
			unpackS.isSlash = true
		}
	}

	unpackS.pastLetter = currentRune
	unpackS.isPastLetter = true
	return nil
}
