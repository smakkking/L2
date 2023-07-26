package dev02

import (
	"errors"
	"unicode"
)

var (
	IncorrectStringError = errors.New("incorrect string")
)

func UnpackString(str []rune) ([]rune, error) {
	var result []rune

	var ecran bool = false
	var symb rune
	var count int = 0

	for _, ch := range str {
		if ecran {
			symb = ch
			ecran = false
		} else {
			if unicode.IsDigit(ch) {
				count += 10*count + int(ch) - 48
			} else {
				if count == 0 && symb != 0 {
					result = append(result, symb)
				} else {
					for i := 0; i < count; i++ {
						result = append(result, symb)
					}
				}
				symb = 0
				count = 0

				if ch == '\\' { // если начинается экранирование, то след символ должен быть 1
					ecran = true
				} else {
					symb = ch
				}
			}

		}
	}

	if (count != 0 && symb == 0) || ecran {
		return result, IncorrectStringError
	}
	if count == 0 && symb != 0 {
		result = append(result, symb)
	} else {
		for i := 0; i < count; i++ {
			result = append(result, symb)
		}
	}

	return result, nil
}
