package transpose

import (
	"regexp"
	"strings"
)

func Transpose(input []string) []string {
	// edge cases
	if len(input) == 0 || len(input[0]) == 0 {
		return nil
	}

	// pad to the left with spaces, dont pad to the right
	max := 0
	// loop through each string
	for i, v := range input {
		// if current string is longer than max
		if len(v) > max {
			// then assign new max
			max = len(v)
			// and loop from the start to current position and add spaces at the end
			for j := 0; j < i; j++ {
				input[j] = input[j] + strings.Repeat(" ", max-len(input[j]))
			}
		}
	}

	// initialize 2d array that equivalent to input size but different dimensions
	bytes := make([][]byte, len(input[0]))
	for i := range bytes {
		// also remember to init the zero-value
		bytes[i] = make([]byte, len(input))
	}

	// loop through the input
	for i := range input {
		for j := range input[i] {
			// and pass to table with reversed dimensions
			bytes[j][i] = input[i][j]
		}
	}

	// convert to string
	var result []string
	// loop through table
	for _, v := range bytes {
		// trim extra allocated bytes at the end
		curr := strings.TrimRight(string(v), "\x00")
		// replace the bytes in the middle of the string to empty space
		curr = regexp.MustCompile(`\x00`).ReplaceAllString(curr, " ")
		// append to result
		result = append(result, curr)
	}

	return result
}
