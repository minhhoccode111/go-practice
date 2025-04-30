package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	primes := []int{}

	floatN := float64(n)
	upper := int(floatN*math.Log2(floatN) + floatN*math.Log2(math.Log2(floatN)))

	// for small prime numbers
	if n < 3 {
		upper = 4
	}

	// create a slice of numbers to upper bound
	for i := 2; i < upper; i++ {
		primes = append(primes, i)
	}

	for i := 0; i < len(primes); i++ {
		curr := primes[i]
		if curr == 0 {
			continue
		}
		// multiply itself and remove. E.g., start with 2, remove 4, 6, 8, 10 etc
		for j := i + curr; j < len(primes); j += curr {
			primes[j] = 0
		}
	}

	count := 0
	// count to nth prime number and return
	for i := 0; i < len(primes); i++ {
		curr := primes[i]
		if curr != 0 {
			count++
		}
		if count == n {
			return curr, nil
		}
	}

	// error if not found
	return 0, errors.New("invalid input")
}
