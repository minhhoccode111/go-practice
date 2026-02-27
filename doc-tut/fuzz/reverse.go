package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(str string) (string, error) {
	if !utf8.ValidString(str) {
		return str, errors.New("input is not valid UTF-8")
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"

	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("Reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
