package main

import "fmt"

func great(name string) string {
	return "Hello, " + name
}

func main() {
	// here we call the function
	msg := great("minhhoccode111")
	fmt.Println(msg)
}
