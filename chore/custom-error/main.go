package main

import "fmt"

// the struc to hold the error
type MyError struct {
	Code    int
	Message string
}

// The custom error method. Must implement this method to satisfy error interface
func (e MyError) Error() string {
	return fmt.Sprintf("Error: %d: %s", e.Code, e.Message)
}

func mightFail(flag bool) error {
	if flag {
		return MyError{Code: 500, Message: "Internal error occurred"}
	}
	return nil
}

func main() {
	err := mightFail(true)
	if err != nil {
		if myError, ok := err.(MyError); ok {
			fmt.Printf("Handled custom error with code %d and message: %s\n", myError.Code, myError.Message)
		} else {
			fmt.Println("An error occurred: ", err)
		}
	}
}
