package allyourbase

import (
	"errors"
	"math"
)

/*
Example the number 42

Base 10:
(4 × 10¹) + (2 × 10⁰)
=> 42

Base 2:
(1 × 2⁵) + (0 × 2⁴) + (1 × 2³) + (0 × 2²) + (1 × 2¹) + (0 × 2⁰) =>
=> 101010

Base 3:
(1 × 3³) + (1 × 3²) + (2 × 3¹) + (0 × 3⁰)
=> 1120
*/

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if !every(inputDigits, func(n int) bool { return n >= 0 && n < inputBase }) {
		return nil, errors.New("all digits must satisfy 0 <= d < input base")
	}
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	inputBaseFloat := float64(inputBase)

	sum := 0
	for i, v := range inputDigits {
		if v < 0 || v >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		powFloat := float64(len(inputDigits) - 1 - i)
		curr := v * int(math.Pow(inputBaseFloat, powFloat))
		sum += curr
	}

	largestPow := float64(0)
	for {
		n := int(math.Pow(float64(outputBase), largestPow))
		if n > sum {
			if largestPow > 0 {
				largestPow--
			}
			break
		}
		largestPow++
	}

	return recursiveOutputBase([]int{}, sum, largestPow, float64(outputBase)), nil
}

func recursiveOutputBase(s []int, num int, pow float64, base float64) []int {
	curr := num / int(math.Pow(base, pow))
	remain := num % int(math.Pow(base, pow))
	s = append(s, curr)
	if pow < 1 {
		return s
	}
	return recursiveOutputBase(s, remain, pow-1, base)
}

// every works like Array.every in JavaScript
func every[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
