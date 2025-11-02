package main

import (
	"fmt"
	"os"
)

// problem:
// create a file
// write to a file
// close a file (with defer)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(s string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing file")
	_, err := fmt.Fprintf(f, "Hello, World!\n")
	if err != nil {
		panic(err)
	}
}
