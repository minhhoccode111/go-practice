package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(s string) string {
	words := strings.Fields(s)
	word := func(s string) string {
		reNotStartWithVowels := regexp.MustCompile(`^([^aeiou]+)`)
		reRuleOne := regexp.MustCompile(`^(xr|yt)`)
		if !reNotStartWithVowels.MatchString(s) || reRuleOne.MatchString(s) {
			return s + "ay"
		}
		reRuleThree := regexp.MustCompile(`^([^aeiou]*qu)`)
		if reRuleThree.MatchString(s) {
			prefix := reRuleThree.FindStringSubmatch(s)
			s = reRuleThree.ReplaceAllString(s, "") + prefix[0]
			return s + "ay"
		}
		reRuleFour := regexp.MustCompile(`^([^aeiou]+)y`)
		if reRuleFour.MatchString(s) {
			prefix := reRuleFour.FindStringSubmatch(s)
			s = "y" + reRuleFour.ReplaceAllString(s, "") + prefix[1]
			return s + "ay"
		}
		prefix := reNotStartWithVowels.FindStringSubmatch(s)
		s = reNotStartWithVowels.ReplaceAllString(s, "") + prefix[0]
		return s + "ay"
	}
	for i, v := range words {
		words[i] = word(v)
	}
	return strings.Join(words, " ")
}
