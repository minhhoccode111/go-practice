package main

import (
	"fmt"
	"math"
)

/*
Interfaces
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22. Vertex (the value type)
doesn't implement Abser because the Abs method is defined only on *Vertex (the
pointer type).
*/

type MyFloat float64

func (m MyFloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	xPow := math.Pow(v.X, 2)
	yPow := math.Pow(v.Y, 2)
	return math.Sqrt(xPow + yPow)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	// in the following line, v is a Vertex (not *Vertex) and does NOT
	// implement Abser.
	// a = v // error
	fmt.Println(a.Abs()) // 5
}
