package main

import (
	_ "fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the
string s. The wc.Test function runs a test suite against the provided function
and prints success or failure.

You might find 'strings.Fields' helpful.
*/

func WordCount(s string) map[string]int {
	var strArr []string
	strArr = strings.Fields(s)

	var result map[string]int
	result = make(map[string]int)

	for _, value := range strArr {
		// fmt.Println(index, value)
		// fmt.Println(result[value])
		if count, ok := result[value]; ok {
			result[value] = count + 1
		} else {
			result[value] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
