package main

import "fmt"

// Embedding
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close
type ReadWriter interface {
	Reader // includes methods of Reader in ReadWriter's method set
	Writer // includes methods of Writer in ReadWriter's method set
}

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "make a sound")
}

func (a Animal) Walk() {
	fmt.Println(a.Name, "walks")
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() { // override
	fmt.Println(d.Name, "barks")
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	d.Walk()         // access animal's Walk method directly without specifying parent
	d.Speak()        // Buddy barks (override)
	d.Animal.Speak() // Buddy makes a sound (call parent)
}
