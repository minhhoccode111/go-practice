package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid floating point number")
		return
	}

	truncated := int(f)
	fmt.Println(truncated)
}
