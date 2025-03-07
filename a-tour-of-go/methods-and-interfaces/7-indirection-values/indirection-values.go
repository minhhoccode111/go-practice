package main

import (
	"fmt"
	"math"
)

/*
Methods and pointer indirection (2)
The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!

while methods with value receivers take either a value or a pointer as the
receiver when they are called:

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK

In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())    // 5
	fmt.Println(AbsFunc(v)) // 5

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())     // 5 - method can be called with pointer (auto dereference)
	fmt.Println(AbsFunc(*p)) // 5 - function have to dereference pointer manually
}
