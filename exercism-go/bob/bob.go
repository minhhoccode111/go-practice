// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

/*

- "Sure."
  This is his response if you ask him a question, such as "How are you?"
  The convention used for questions is that it ends with a question mark.
- "Whoa, chill out!"
  This is his answer if you YELL AT HIM.
  The convention used for yelling is ALL CAPITAL LETTERS.
- "Calm down, I know what I'm doing!"
  This is what he says if you yell a question at him.
- "Fine. Be that way!"
  This is how he responds to silence.
  The convention used for silence is nothing, or various combinations of whitespace characters.
- "Whatever."
  This is what he answers to anything else.

*/

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	isUpper := strings.ToUpper(remark) == remark
	reHasLetters := regexp.MustCompile(`[A-Za-z]`)
	reQuestion := regexp.MustCompile(`\?$`)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if isUpper && reQuestion.MatchString(remark) && reHasLetters.MatchString(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if reQuestion.MatchString(remark) {
		return "Sure."
	}

	if isUpper && reHasLetters.MatchString(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
