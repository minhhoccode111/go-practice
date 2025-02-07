package main

import "fmt"

/*
Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying
array.

Changing the elements of a slice modifies the corresponding elements of its
underlying array.

Other slices that share the same underlying array will see those changes.
*/

func main() {
	printSlicesPointersExample()
}

func printSlicesPointersExample() {
	strings := [4]string{"Hello", "World", "From", "Go"}
	fmt.Println(strings) // [Hello World From Go]

	a := strings[0:2]
	b := strings[1:3]
	fmt.Println(a, b) // [Hello World] [World From]

	b[0] = "XXX"
	fmt.Println(a, b)    // [Hello XXX] [XXX From]
	fmt.Println(strings) // [Hello XXX From Go]
}
