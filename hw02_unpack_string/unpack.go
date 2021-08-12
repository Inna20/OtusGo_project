package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var newString strings.Builder
	var curRune rune
	var result string
	var resultError error

	for _, symbol := range str {
		if unicode.IsDigit(symbol) {
			if curRune == 0 {
				resultError = ErrInvalidString
				break
			}

			var count, _ = strconv.Atoi(string(symbol))
			if count > 0 {
				newString.WriteString(strings.Repeat(string(curRune), count))
			}
			// пустая руна, чтобы в следующей итерации она не была записана в результат
			curRune = rune(0)
		} else {
			if curRune != 0 {
				newString.WriteString(string(curRune))
			}
			curRune = symbol
		}
	}

	if resultError == nil {
		// Если последний символ не цифра - тоже запишем в строку
		if !unicode.IsDigit(curRune) && curRune != 0 {
			newString.WriteString(string(curRune))
		}

		result = newString.String()
	}

	return result, resultError
}
