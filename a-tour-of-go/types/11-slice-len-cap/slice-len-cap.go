package main

import "fmt"

/*
Slice length and capacity
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array,
counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions
len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient
capacity. Try changing one of the slice operations in the example program to
extend it beyond its capacity and see what happens.
*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
	// s[2] = 11 // error
	// s[3] = 13 // error
	// fmt.Println(s[2], s[3]) // error

	// Restore its length.
	s = s[:cap(s)]
	printSlice(s) // len=4 cap=4 [5 7 11 13]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
