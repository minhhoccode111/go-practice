package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	// ignore cases
	input = strings.ToLower(input)

	// remove not word characters
	re := regexp.MustCompile(`\W`)
	input = re.ReplaceAllString(input, "")

	// create a dictionary
	dict := map[rune]int{}
	for r := 'a'; r <= 'z'; r++ {
		dict[r] = 0
	}
	// loop through input to count each letter's appearance
	for _, v := range input {
		dict[v]++
	}

	// loop through dict to see if each letter is present
	for _, v := range dict {
		if v < 1 {
			return false
		}
	}

	return true
}
