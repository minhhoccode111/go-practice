package cryptosquare

import (
	"regexp"
	"strings"
)

func Encode(pt string) string {
	// normalized the input
	re := regexp.MustCompile(`[^a-z0-9]`)
	pt = re.ReplaceAllString(strings.ToLower(pt), "")

	// init values for columns, rows and length
	c, r, l := 0, 0, 0

	// while l is less than the length of the input
	for l < len(pt) {
		// r must be greater than or equal to c
		if c == r {
			r++
		} else {
			c++
		}
		// increase new length base on new rows and columns
		l = c * r
	}

	// padEnd the input
	for len(pt) != l {
		pt = pt + " "
	}

	table := make([]string, r)

	// divide the input into columns
	for i := 0; i < len(pt); i++ {
		col := i % len(table)
		char := string(pt[i])
		table[col] += char
	}

	// join to single line separate each row by a space
	return strings.Join(table, " ")
}
