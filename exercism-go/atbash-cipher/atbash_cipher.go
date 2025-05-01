package atbash

import (
	"regexp"
	"strings"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[\s,.]`).ReplaceAllString(s, "")
	dict := map[rune]rune{}
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	for i, v := range plain {
		dict[v] = rune(cipher[i])
	}
	result := []rune{}
	count := 0
	for _, v := range s {
		if count == 5 {
			result = append(result, ' ')
			count = 0
		}
		// case valid characaters
		if r, ok := dict[v]; ok {
			result = append(result, r)
		} else {
			// case numbers
			result = append(result, v)
		}
		count++
	}
	return string(result)
}
