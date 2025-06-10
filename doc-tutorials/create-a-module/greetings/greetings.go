package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was given, return a value that embeds the name in a greeting
	// message
	// message := fmt.Sprint(randomFormat()) // break the greetings.Hello() function to view a failing test
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting
// message.
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages.
	messages := make(map[string]string)
	// loop through the received slice of names, calling the Hello function to
	// get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns on of a set of greeting messages the returned message
// is selected at random.
func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// return a randomly selected message format by specifying a random index
	// for the slice of formats
	return formats[rand.Intn(len(formats))]
}
