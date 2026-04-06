package kata

import (
	"sort"
)

func ExpressionMatter(a int, b int, c int) int {
	s := []int{
		a + b + c,
		a + (b * c),
		(a + b) * c,
		(a * b) + c,
		a * (b + c),
		a * b * c,
	}
	sort.Ints(s)
	return s[len(s)-1]
}
