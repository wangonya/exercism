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

// using strings.Count
// func IsIsogram(phrase string) bool {
// 	phrase = strings.ToLower(phrase)
// 	for _, chr := range phrase {
// 		if !unicode.IsLetter(chr) {
// 			continue
// 		}
// 		if strings.Count(phrase, string(chr)) > 1 {
// 			return false
// 		}
// 	}
// 	return true
// }

// using Bitfield
// func IsIsogram(phrase string) bool {
// 	theEnd := len(phrase)
// 	var letterFlags uint32 = 0

// 	for i := 0; i < theEnd; i++ {
// 		letter := phrase[i]
// 		if letter > 96 && letter < 123 {
// 			if (letterFlags & (1 << (letter - 'a'))) != 0 {
// 				return false
// 			} else {
// 				letterFlags |= (1 << (letter - 'a'))
// 			}
// 		} else if letter > 64 && letter < 91 {
// 			if (letterFlags & (1 << (letter - 'A'))) != 0 {
// 				return false
// 			} else {
// 				letterFlags |= (1 << (letter - 'A'))
// 			}
// 		}
// 	}
// 	return true
// }
