package main

import "fmt"

/*
Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and
many others) look for this interface to print values.
*/

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"ASD LKJ", 123}
	b := Person{"QWE OIU", 987}
	fmt.Println(&a, &b) // ASD LKJ (123 years) QWE OIU (987 years)
}
