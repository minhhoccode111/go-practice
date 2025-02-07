package main

import (
	"fmt"
	"math"
)

/*
Pointer receivers
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also,
T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

Methods with pointer receivers can modify the value to which the receiver
points (as Scale does here). Since methods often need to modify their receiver,
pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and
observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original
Vertex value. (This is the same behavior as for any other function argument.)
The Scale method must have a pointer receiver to change the Vertex value
declared in the main function.
*/

type Point struct {
	X, Y float64
}

// this will pass a pointer to 'p'
func (p *Point) Scale(time float64) {
	p.X = p.X * time
	p.Y = p.Y * time
}

// this will pass a copy of 'p'
func (p Point) Abs() float64 {
	xPow := math.Pow(p.X, 2)
	yPow := math.Pow(p.Y, 2)
	return math.Sqrt(xPow + yPow)
}

func main() {
	var v Point
	v = Point{math.SqrtPi, math.Pi}
	fmt.Println(v)
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}
