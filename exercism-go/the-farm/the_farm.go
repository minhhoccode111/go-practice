package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalAmount, totalAmountErr := fc.FodderAmount(numCows)
	if totalAmountErr != nil {
		return 0.0, totalAmountErr
	}
	factor, factorErr := fc.FatteningFactor()
	if factorErr != nil {
		return 0.0, factorErr
	}
	return totalAmount * factor / float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

type InvalidCowsError struct {
	numCows int
	msg     string
}

func (ic *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", ic.numCows, ic.msg)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{numCows: numCows, msg: "there are no negative cows"}
	}
	if numCows == 0 {
		return &InvalidCowsError{numCows: numCows, msg: "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
