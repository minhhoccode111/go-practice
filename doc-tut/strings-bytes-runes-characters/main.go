package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// sample := []byte{0xbd, 0xb2, 0x3d, 0xbc, 0x20, 0xe2, 0x8c, 0x98}

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := range len(sample) {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()
	for i := range len(sample) {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	fmt.Println()
	fmt.Println()
	fmt.Println("stringLiteral")
	stringLiteral()

	fmt.Println()
	fmt.Println()
	fmt.Println("rangeLoop")
	rangeLoop()

	fmt.Println()
	fmt.Println()
	fmt.Println("loopUtf8")
	loopUtf8()

}

func loopUtf8() {
	const nihongo = "日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func rangeLoop() {
	const nihongo = "日U+FFFF本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func stringLiteral() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := range len(placeOfInterest) {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}
