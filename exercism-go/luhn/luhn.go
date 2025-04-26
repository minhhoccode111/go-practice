package luhn

import (
	"regexp"
	"strings"
)

func Valid(id string) bool {
	// remove spaces
	id = strings.ReplaceAll(id, " ", "")
	// cover edge cases
	re := regexp.MustCompile(`\D`)
	if len(id) <= 1 || re.MatchString(id) {
		return false
	}
	// reverse the string
	runes := []rune(id)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// calculate the sum
	sum := 0
	for i, v := range runes {
		digit := int(v - '0')
		// double second digit from the right (after reverse)
		if i%2 == 1 {
			digit = digit * 2
			// if greater than 9 subtract 9
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
