package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, v := range s {
		currInt, _ := strconv.Atoi(string(v))
		currPow := math.Pow(float64(currInt), float64(len(s)))
		sum += int(currPow)
	}
	return sum == n
}
