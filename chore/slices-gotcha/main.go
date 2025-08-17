package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
A possible "gotcha"
- Re-slicing a slice doesn't make a copy of the underlying array. The full array
will be kept in memory until it is no longer referenced. Occasionally this can
cause the program to hold all the data in memory when only a small piece of it
is needed
*/

var digitRegexp = regexp.MustCompile("[0-9]+")

// FindDigits function loads a file into memory and searches it for the first
// group of consecutive digits, returning them as new slice
// This code behaves as advertised but the returned `[]byte` points to an array
// containing the entire file. Since the slice references the original array, as
// long as the slice is kept around the garbage collector can't release the array
// the few useful bytes of the file keep the entire contents in memory
func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return digitRegexp.Find(b)
}

// CopyDigits to fix this problem one can copy the interesting data to a new
// slice before returning it
func CopyDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// AppendDigits or a more concise version of previous function could be
// constructed by using `append`
func AppendDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := append([]byte{}, b...)
	return c
}

func main() {
	fmt.Println("Hello, World!")
}
