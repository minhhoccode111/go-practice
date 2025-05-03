package sieve

func Sieve(n int) []int {
	// cover edge cases
	if n < 2 {
		return nil
	}

	primes := []int{}

	// fill the slice to the upper bound
	for i := 2; i <= n; i++ {
		primes = append(primes, i)
	}

	// loop through the slice and mark non-prime numbers
	for i, v := range primes {
		// skip if already marked
		if v == 0 {
			continue
		}

		// e.g.,
		// 2 will mark 4, 6, 8, 10, etc. as non-prime
		// 3 will mark 9, 12, 15, 18, etc. as non-prime
		// 5 will mark 25, 30, 35, 40, etc. as non-prime
		for j := i + v; j < len(primes); j += v {
			primes[j] = 0
		}
	}

	// filter to remove zeros from the slice and return
	return filter(primes, func(e int) bool {
		return e != 0
	})
}

// filter works like JavaScript's Array.filter method
func filter[T any](slice []T, callback func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
