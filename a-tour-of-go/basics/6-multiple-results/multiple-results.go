package main

import "fmt"

/*
Multiple results
A function can return any number of results.

The swap function returns two strings.
*/

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
