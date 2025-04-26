package scrabble

import "strings"

func scoreChar(c rune) int {
	switch {
	case c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'l' || c == 'n' || c == 'r' || c == 's' || c == 't':
		return 1
	case c == 'd' || c == 'g':
		return 2
	case c == 'b' || c == 'c' || c == 'm' || c == 'p':
		return 3
	case c == 'f' || c == 'h' || c == 'v' || c == 'w' || c == 'y':
		return 4
	case c == 'k':
		return 5
	case c == 'j' || c == 'x':
		return 8
	case c == 'q' || c == 'z':
		return 10
	default:
		panic("character not allowed")
	}
}

func Score(word string) int {
	word = strings.ToLower(word)
	count := 0
	for _, v := range word {
		count += scoreChar(v)
	}
	return count
}
