package isbn

import "strings"

func IsValidISBN(in string) bool {
	in = strings.ReplaceAll(in, "-", "")
	if len(in) != 10 {
		return false
	}
	sum := 0
	for i, v := range in {
		if v == 'X' {
			sum += 10
			continue
		}
		n := int(v - '0')
		sum += n * (10 - i)
	}
	return sum%11 == 0
}
