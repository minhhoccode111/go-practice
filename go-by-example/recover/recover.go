package main

import "fmt"

func mayPanic() {
	panic("Hello, World!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered:", r)
		}
	}()

	mayPanic()

	fmt.Println("doesn't reach here")
}
