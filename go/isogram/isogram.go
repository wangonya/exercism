package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(word, "-", ""), " ", ""))
	letters := map[rune]rune{}

	for _, r := range word {
		if _, ok := letters[r]; ok {
			return false
		} else {
			letters[r] = r
		}
	}
	return true
}
