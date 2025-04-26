package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	dict := map[rune]bool{}
	for _, v := range word {
		if v == rune('-') || v == rune(' ') {
			continue
		}
		_, exists := dict[v]
		if exists {
			return false
		}
		dict[v] = true
	}
	return true
}
