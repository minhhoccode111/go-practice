package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range []rune(log) {
		switch {
		case v == '❗':
			return "recommendation"
		case v == '🔍':
			return "search"
		case v == '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeStr := []rune(log)
	for i, v := range runeStr {
		if v == oldRune {
			runeStr[i] = newRune
		}
	}
	return string(runeStr)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
