package main

import "fmt"

/*
Defer
A defer statement defers the execution of a function until the surrounding
function returns.

The deferred call's arguments are evaluated immediately, but the function call
is not executed until the surrounding function returns.
*/

func main() {
	printDefer()
	// Hello,
	// World!
}

func printDefer() {
	defer fmt.Println("World!")

	fmt.Println("Hello,")
}
