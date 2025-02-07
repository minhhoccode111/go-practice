package main

import "fmt"

/*
Nil slices
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.
*/

func main() {
	var s []int
	fmt.Printf("len: %d, cap: %d, val: %v", len(s), cap(s), s) // len: 0, cap: 0, val: []nil!
	if s == nil {
		fmt.Println("nil!")
	}
}
