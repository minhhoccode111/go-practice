package main

import "fmt"

/*
Generic types
In addition to generic functions, Go also supports generic types. A type can be
parameterized with a type parameter, which could be useful for implementing
generic data structures.

This example demonstrates a simple type declaration for a singly-linked list
holding any type of value.

As an exercise, add some functionality to this list implementation.
*/

// List represents a singly-linked list that holds
// values of any type.
type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func main() {
	var head Node[int]
	head = Node[int]{0, nil}
	pointer := &head
	for i := 1; i < 10; i++ {
		pointer.Next = &Node[int]{i, nil}
		pointer = pointer.Next
	}

	for p := &head; p != nil; p = p.Next {
		fmt.Print(p.Val, " -> ")
	}
	fmt.Println("nil")

	// output: 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> nil
}
