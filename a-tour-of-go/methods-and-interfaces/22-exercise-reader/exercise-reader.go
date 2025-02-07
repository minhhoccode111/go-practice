package main

import (
	"strings"

	"golang.org/x/tour/reader"
)

/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (int, error) {
	r := strings.NewReader("A")
	n, err := r.Read(b)
	return n, err
}

func main() {
	reader.Validate(MyReader{}) // OK!
}
