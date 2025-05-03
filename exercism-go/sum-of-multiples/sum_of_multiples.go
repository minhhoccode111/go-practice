package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	dict := map[int]int{}
	for _, v := range divisors {
		if v == 0 {
			continue
		}
		tmp := v
		for v < limit {
			dict[v] = 1
			v += tmp
		}
	}
	for k := range dict {
		sum += k
	}
	return sum
}
