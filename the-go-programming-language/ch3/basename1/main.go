// basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package basename1

func basename(s string) string {
	// discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
