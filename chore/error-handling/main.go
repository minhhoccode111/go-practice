package main

import (
	"errors"
	"fmt"
)

func main() {
	result0, err := divide(1, 0)
	if err != nil {
		fmt.Println(err.Error()) // division by zero is not allowed
	} else {
		fmt.Println(result0)
	}

	result1, err := divide(1, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result1) // 0.5
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
