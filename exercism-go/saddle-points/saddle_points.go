package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int

type Pair struct {
	a int
	b int
}

func New(s string) (*Matrix, error) {
	lines := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})

	result := Matrix{}
	for _, line := range lines {
		digitsString := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		})

		digitsInt := []int{}
		for _, digitString := range digitsString {
			digit, err := strconv.Atoi(digitString)
			if err != nil {
				return nil, errors.New("invalid string")
			}
			digitsInt = append(digitsInt, digit)
		}

		result = append(result, digitsInt)
	}

	return &result, nil
}

func (m *Matrix) Saddle() []Pair {
	if len(*m) == 0 {
		return nil
	}
	if len((*m)[0]) == 0 {
		return nil
	}
	var result []Pair
	maxRow := make([]int, len(*m))
	minCol := make([]int, len((*m)[0]))
	for i, r := range *m {
		max := r[0]
		for _, c := range r {
			if c > max {
				max = c
			}
		}
		maxRow[i] = max
	}
	for i := 0; i < len((*m)[0]); i++ {
		min := (*m)[0][i]
		for j := 0; j < len(*m); j++ {
			c := (*m)[j][i]
			if c < min {
				min = c
			}
		}
		minCol[i] = min
	}

	for i, r := range *m {
		for j, c := range r {
			if c == maxRow[i] && c == minCol[j] {
				result = append(result, Pair{i + 1, j + 1})
			}
		}
	}

	return result
}
