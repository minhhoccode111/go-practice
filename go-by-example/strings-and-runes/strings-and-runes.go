package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	fmt.Println("Len:", len(s)) // 18

	for i := range len(s) {
		fmt.Printf("%x ", s[i]) // e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		/*
		   U+0E2A 'ส' starts at 0
		   U+0E27 'ว' starts at 3
		   U+0E31 'ั' starts at 6
		   U+0E2A 'ส' starts at 9
		   U+0E14 'ด' starts at 12
		   U+0E35 'ี' starts at 15
		*/
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
		/*
		   U+0E2A 'ส' starts at 0
		   found so sua
		   U+0E27 'ว' starts at 3
		   U+0E31 'ั' starts at 6
		   U+0E2A 'ส' starts at 9
		   found so sua
		   U+0E14 'ด' starts at 12
		   U+0E35 'ี' starts at 15
		*/
	}
}

func examineRune(r rune) {
	switch r {
	case 't':
		fmt.Println("found tee")
	case 'ส':
		fmt.Println("found so sua")
	}
}
