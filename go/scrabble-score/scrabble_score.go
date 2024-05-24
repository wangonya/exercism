package scrabble

import (
	"unicode"
	"unicode/utf8"
)

func value(r rune) int {
	value := 0
	r = unicode.ToUpper(r)
	switch r {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		value = 1
	case 'D', 'G':
		value = 2
	case 'B', 'C', 'M', 'P':
		value = 3
	case 'F', 'H', 'V', 'W', 'Y':
		value = 4
	case 'K':
		value = 5
	case 'J', 'X':
		value = 8
	case 'Q', 'Z':
		value = 10
	}
	return value
}

func Score(word string) int {
	points := 0
	for len(word) > 0 {
		r, size := utf8.DecodeRuneInString(word)
		points += value(r)

		word = word[size:]
	}
	return points
}
