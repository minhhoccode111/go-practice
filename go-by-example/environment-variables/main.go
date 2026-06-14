package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("foo", "1")
	fmt.Println(os.Getenv("foo"))
	fmt.Println(os.Getenv("bar"))
	fmt.Println()
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
