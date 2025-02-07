package main

import "fmt"

/*
Struct Fields
Struct fields are accessed using a dot.
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 8
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
