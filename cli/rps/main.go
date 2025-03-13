package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	play()
}

func play() {
	printGreeting()          // function
	botValue := rand.Intn(3) // random
	userValue := getUserValue()
	printResult(userValue, botValue)
}

func printGreeting() {
	fmt.Println("Let's play Rock, Paper, Scissors Game!")
	fmt.Println("You choose [R]ock, [P]aper, or [S]cissors? ")
}

func getUserValue() int {
	isValidInput := func(input string) (int, bool) { // anonymous function
		inputMap := map[string]int{"r": 0, "p": 1, "s": 2} // map
		val, exists := inputMap[strings.ToLower(input)]
		return val, exists
	}

	var userInput string

	for {
		_, e := fmt.Scanln(&userInput)                                        // pointer
		if userValue, exists := isValidInput(userInput); e == nil && exists { // short if
			return userValue
		}
		fmt.Println("Please input 1 character: R/r/P/p/S/s")
	}
}

func printResult(userValue int, botValue int) {
	dict := map[int]string{0: "rock", 1: "paper", 2: "scissors"}
	switch { // switch case
	case userValue == 0 && botValue == 1, userValue == 1 && botValue == 2, userValue == 2 && botValue == 0: // multiple case
		fmt.Println("You lose!")
	case botValue == 0 && userValue == 1, botValue == 1 && userValue == 2, botValue == 2 && userValue == 0:
		fmt.Println("You win!")
	default:
		fmt.Println("It's a tie!")
	}
	fmt.Printf("You chose %s, Bot chose %s\n", dict[userValue], dict[botValue]) // string interpolation
}
