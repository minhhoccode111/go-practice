package main

import "fmt"

/*
Type parameters

Go functions can be written to work on multiple types using type parameters.
The type parameters of a function appear between brackets, before the
function's arguments.

func Index[T comparable](s []T, x T) int

This declaration means that s is a slice of any type T that fulfills the
built-in constraint comparable. x is also a value of the same type.

comparable is a useful constraint that makes it possible to use the == and !=
operators on values of the type. In this example, we use it to compare a value
to all slice elements until a match is found. This Index function works for any
type that supports comparison.
*/

func main() {
	// index works on a slice of ints
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(Index(si, 3)) // 2
	// index works on a slice of strings
	ss := []string{"hello", "world", "from", "go"}
	fmt.Println(Index(ss, "mhc")) // -1
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if x == v {
			return i
		}
	}
	return -1
}
