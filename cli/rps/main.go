package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("Let's play Rock, Paper, Scissors Game!")
	fmt.Println("You choose [R]ock, [P]aper, or [S]cissors? ")
	botValue := rand.Intn(3)
	var userInput string
	var userValue int
	getUserInput(&userInput, &userValue)
	printResult(userValue, botValue)
}

func getUserInput(userInput *string, userValue *int) {
	for {
		_, e := fmt.Scanln(userInput)
		if val, exists := isValidInput(*userInput); e == nil && exists {
			*userValue = val
			break
		}
		fmt.Println("Please input 1 character: R/r/P/p/S/s")
	}
}

func isValidInput(input string) (int, bool) {
	inputMap := map[string]int{
		"r": 0,
		"p": 1,
		"s": 2,
	}
	val, exists := inputMap[strings.ToLower(input)]
	return val, exists
}

func printResult(userValue int, botValue int) {
	dict := map[int]string{0: "rock", 1: "paper", 2: "scissors"}
	switch {
	case userValue == 0 && botValue == 1, userValue == 1 && botValue == 2, userValue == 2 && botValue == 0:
		fmt.Println("You lose!")
	case botValue == 0 && userValue == 1, botValue == 1 && userValue == 2, botValue == 2 && userValue == 0:
		fmt.Println("You win!")
	default:
		fmt.Println("It's a tie!")
	}
	fmt.Printf("You chose %s, Bot chose %s\n", dict[userValue], dict[botValue])
}
