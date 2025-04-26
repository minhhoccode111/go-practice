package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n cannot be less than 1")
	}
	count := 0
	for n != 1 {
		if n%2 == 0 {
			count++
			n /= 2
		} else {
			count++
			n = n*3 + 1
		}
	}
	return count, nil
}
