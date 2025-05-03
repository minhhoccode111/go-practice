package prime

func Factors(n int64) []int64 {
	result := []int64{}
	curr := int64(2)
	for n > 1 {
		if n%curr == 0 {
			n /= curr
			result = append(result, curr)
		} else {
			curr++
		}
	}
	return result
}
