package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname, lname [20]rune
}

func New(fname, lname [20]rune) Name {
	return Name{fname, lname}
}

func ConvertToFixedTwenty(s string) [20]rune {
	var arr [20]rune
	runes := []rune(s)
	copy(arr[:], runes)
	return arr
}

func main() {
	fmt.Print("Input file name: ")
	scannerStdin := bufio.NewScanner(os.Stdin)
	scannerStdin.Scan()
	fileName := scannerStdin.Text()

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var names []Name

	scannerFile := bufio.NewScanner(file)
	for scannerFile.Scan() {
		parts := strings.Fields(scannerFile.Text())
		if len(parts) < 2 {
			continue
		}
		names = append(names, New(
			ConvertToFixedTwenty(parts[0]),
			ConvertToFixedTwenty(parts[1]),
		))
	}

	for _, v := range names {
		fmt.Printf("First name: %s - Last name: %s\n", string(v.fname[:]), string(v.lname[:]))
	}
}
