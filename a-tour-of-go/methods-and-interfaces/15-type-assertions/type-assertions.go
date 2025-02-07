package main

import "fmt"

/*
Type assertions
A type assertion provides access to an interface value's underlying concrete
value.

t := i.(T)

This statement asserts that the interface value i holds the concrete type T and
assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can
return two values: the underlying value and a boolean value that reports
whether the assertion succeeded.

t, ok := i.(T)

If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic
occurs.

Note the similarity between this syntax and that of reading from a map.
*/

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	// f = i.(float64) // panic
	fmt.Println(f) // 0
}
