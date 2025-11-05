package palindrome

import (
	"fmt"
	"strconv"
)

// Define Product type here.
type Product struct {
	N              int
	Factorizations [][2]int
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	return s == reverse(s)
}

func sortFactors(i, j int) [2]int {
	if i > j {
		return [2]int{j, i}
	}
	return [2]int{i, j}
}

func errorResponse(s string) (Product, Product, error) {
	return Product{}, Product{}, fmt.Errorf(s)
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return errorResponse("fmin > fmax")
	}

	minP := Product{N: fmin * fmax, Factorizations: make([][2]int, 0)}
	maxP := Product{N: 0, Factorizations: make([][2]int, 0)}
	var minTable map[[2]int]bool
	var maxTable map[[2]int]bool

	for i := fmin; i <= fmax; i++ {
		for j := fmin; j <= fmax; j++ {
			f := sortFactors(i, j)

			n := i * j
			if !isPalindrome(n) {
				continue
			}

			// fmt.Println(n)

			if n < minP.N {
				minP.N = n
				minP.Factorizations = make([][2]int, 0)
				minP.Factorizations = append(minP.Factorizations, f)
				minTable = map[[2]int]bool{}
				minTable[f] = true
			}

			if _, ok := minTable[f]; n == minP.N && !ok {
				minP.Factorizations = append(minP.Factorizations, f)
			}

			if n > maxP.N {
				maxP.N = n
				maxP.Factorizations = make([][2]int, 0)
				maxP.Factorizations = append(maxP.Factorizations, f)
				maxTable = map[[2]int]bool{}
				maxTable[f] = true
			}

			if _, ok := maxTable[f]; n == maxP.N && !ok {
				maxP.Factorizations = append(maxP.Factorizations, f)
			}
		}
	}

	if maxP.N == 0 {
		return errorResponse("no palindromes")
	}

	return minP, maxP, nil
}
