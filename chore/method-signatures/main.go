package main

import "fmt"

type Greeter interface {
	Greet() string
}

// Person struct implements the Greeter interface
type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, World! From " + p.Name + "!"
}

// SayHello accepts any type that satisfies the interface value Greeter.
func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	p := Person{Name: "minhhoccode111"}
	SayHello(p)
}
