package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(in string) Frequency {
	in = strings.ToLower(in)
	re := regexp.MustCompile(`\w+'\w+|\w+`)
	words := re.FindAllString(in, -1)
	dict := Frequency{}
	for _, word := range words {
		dict[word]++
	}
	return dict
}
