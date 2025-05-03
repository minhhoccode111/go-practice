package phonenumber

import (
	"errors"
	"regexp"
	"strings"
)

var digitsRe = regexp.MustCompile(`\d+`)
var validRe = regexp.MustCompile(`^1*([2-9]\d{2})([2-9]\d{2})(\d{4})$`)

func Number(input string) (string, error) {
	input = strings.Join(digitsRe.FindAllString(input, -1), "")
	if !validRe.MatchString(input) {
		return "", errors.New("invalid input")
	}
	return validRe.ReplaceAllString(input, "$1$2$3"), nil
}

func AreaCode(input string) (string, error) {
	input = strings.Join(digitsRe.FindAllString(input, -1), "")
	if !validRe.MatchString(input) {
		return "", errors.New("invalid input")
	}
	return validRe.ReplaceAllString(input, "$1"), nil
}

func Format(input string) (string, error) {
	input = strings.Join(digitsRe.FindAllString(input, -1), "")
	if !validRe.MatchString(input) {
		return "", errors.New("invalid input")
	}
	return validRe.ReplaceAllString(input, "($1) $2-$3"), nil
}
