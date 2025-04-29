package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(in string) Frequency {
	in = strings.ToLower(in)
	reSeparator := regexp.MustCompile(`[.,'":!?\t\n\s&@$%^&]`)
	dict := Frequency{}
	currWord := ""

	for i, v := range in {
		// if is not separator or is valid single quote
		if s := string(v); !reSeparator.MatchString(s) || isValidSingleQuote(i, in) {
			currWord += s
			continue
		}

		// not a word & current work not empty
		if currWord != "" {
			dict[currWord]++
			currWord = ""
		}
	}
	if currWord != "" {
		dict[currWord]++
		currWord = ""
	}

	return dict
}

func isValidSingleQuote(i int, s string) bool {
	r := s[i]

	// if is not single quote or if index out of bounds
	if r != '\'' || i < 1 || i > len(s)-2 {
		return false
	}

	// match word characters
	re := regexp.MustCompile(`\w`)

	// assume index is not out of bounds
	before := s[i-1]
	after := s[i+1]

	// before and after must be word characters
	return re.MatchString(string(before)) && re.MatchString(string(after))
}
