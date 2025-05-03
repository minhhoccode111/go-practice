package series

func All(n int, s string) []string {
	var result []string
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
