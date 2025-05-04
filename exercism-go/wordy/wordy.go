package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(q string) (int, bool) {
	words := strings.Fields(q)
	if words[0] != "What" || words[1] != "is" || !regexp.MustCompile(`\d+\?$`).MatchString(words[len(words)-1]) {
		return 0, false
	}
	// remove prefix
	words = words[2:]

	expect := "num"
	prevNum := 0
	prevOpt := "plus"
	reNum := regexp.MustCompile(`-{0,1}\d+`)

	for _, word := range words {
		word = strings.ReplaceAll(word, "?", "")
		if word == "by" {
			continue
		}
		if reNum.MatchString(word) && expect == "num" {
			currNum, _ := strconv.Atoi(word)

			prevNum = cal(prevNum, currNum, prevOpt)
			expect = "opt"
		} else if isValidOpt(word) && expect == "opt" {
			prevOpt = word
			expect = "num"
		} else {
			return 0, false
		}
	}
	return prevNum, true
}

func isValidOpt(opt string) bool {
	return opt == "plus" || opt == "minus" || opt == "divided" || opt == "multiplied"
}

func cal(a, b int, opt string) int {
	switch opt {
	case "plus":
		return a + b
	case "minus":
		return a - b
	case "divided":
		return a / b
	case "multiplied":
		return a * b
	default:
		panic("unknown opt")
	}
}
