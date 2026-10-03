package main

import "strings"

func main() {}

func CommaFloatAndSign(s string) string {
	sign := ""
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		sign, s = s[:1], s[1:]
	}
	intPart, fracPart := s, ""
	if i := strings.IndexByte(s, '.'); i >= 0 {
		intPart, fracPart = s[:i], s[i:]
	}
	return sign + comma(intPart) + fracPart
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	first := n % 3
	if first == 0 {
		first = 3
	}
	var b strings.Builder
	b.WriteString(s[:first])
	for i := first; i < n; i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
