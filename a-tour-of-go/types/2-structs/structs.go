package main

import "fmt"

/*
Struct
A struct is a collection of fields.
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	tmp := Vertex{1, 2}
	fmt.Println(tmp)   // {1 2}
	fmt.Println(tmp.X) // 1
	fmt.Println(tmp.Y) // 2
}
