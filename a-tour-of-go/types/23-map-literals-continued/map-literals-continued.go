package main

import "fmt"

/*
Map literals continued
If the top-level type is just a type name, you can omit it from the elements of
the literal.
*/

var m = map[string]struct {
	Lat, Long float64
}{
	"asd": {12.12, 34.34},
	"qwe": {34.34, 12.12},
}

func main() {
	fmt.Println(m) // map[asd:{12.12 34.34} qwe:{34.34 12.12}]
}
