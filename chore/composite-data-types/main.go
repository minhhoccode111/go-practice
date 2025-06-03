package main

import "fmt"

// define a composite data type using a struct
type Person struct {
	name string
	age  int
}

func main() {
	// create a new struct and print key value pairs
	person := Person{name: "minhhoccode111", age: 18}
	fmt.Println("Name: ", person.name, "-", "Age: ", person.age)
}
