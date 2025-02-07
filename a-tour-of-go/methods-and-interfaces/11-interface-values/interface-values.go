package main

import (
	"fmt"
)

/*
Interface values
Under the hood, interface values can be thought of as a tuple of a value and a
concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on
its underlying type.
*/

type IHuman interface {
	SaySomething()
}

type Human struct {
	SomethingToSay string
}

func (h *Human) SaySomething() {
	fmt.Println(h.SomethingToSay)
}

type Floating float64

func (f Floating) SaySomething() {
	fmt.Println(f)
}

func describe(i IHuman) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func main() {
	var i IHuman

	i = &Human{"Hello"}
	describe(i)      // (&{Hello}, *main.Human)
	i.SaySomething() // Hello

	i = Floating(3.14)
	describe(i)      // (3.14, main.Floating)
	i.SaySomething() // 3.14
}
