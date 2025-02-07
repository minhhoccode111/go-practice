package main

import "fmt"

/*
Maps
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for
use.
*/

type V struct {
	Lat, Long float64
}

var m map[string]V

func main() {
	m = make(map[string]V)
	m["Hello World"] = V{
		3.14, -3.14,
	}
	fmt.Println(m["Hello World"]) // {3.14 -3.14}
}
