package encode

import (
	"regexp"
	"strconv"
)

func updateResultEncode(curr rune, occurence int, result string) string {
	if curr != 0 {
		if occurence > 1 {
			result += strconv.Itoa(occurence)
		}
		result += string(curr)
	}
	return result
}
func RunLengthEncode(input string) string {
	result := ""
	var curr rune
	occurence := 0
	for _, v := range input {
		if curr == v {
			occurence++
			continue
		}
		result = updateResultEncode(curr, occurence, result)
		curr = v
		occurence = 1
	}
	result = updateResultEncode(curr, occurence, result)
	return result
}
func RunLengthDecode(input string) string {
	result := ""
	re := regexp.MustCompile(`(\d+)?([a-zA-Z\s])`)
	submatches := re.FindAllStringSubmatch(input, -1)
	for _, v := range submatches {
		if v[1] == "" {
			result += v[2]
			continue
		}
		n, _ := strconv.Atoi(v[1])
		for i := 0; i < n; i++ {
			result += v[2]
		}
	}
	return result
}
