package main

import "fmt"

func main() {
	fmt.Println(comma("12345678"))
	/*
	   $ go run comma.go
	   12,345,678
	*/
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
