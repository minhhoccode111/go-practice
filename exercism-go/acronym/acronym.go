package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re := regexp.MustCompile(`[a-zA-Z]+'[a-zA-Z]+|[a-zA-Z]+`)
	matches := re.FindAllString(s, -1)
	result := ""
	for _, v := range matches {
		result += strings.ToUpper(string(v[0]))
	}
	return result
}
