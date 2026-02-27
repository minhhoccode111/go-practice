package main

import (
	"doc/tutorial/greetings"
	"fmt"
	"log"
)

func main() {
	// set properties of the predefined Logger, including the log entry prefix
	// and a flag to disable printing the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"TypeScript", "Golang", "JavaScript", "C"}

	// request a greeting messages for the names.
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message to the console
	fmt.Println(message)
}
