package lsproduct

import (
	"errors"
	_ "fmt"
	"math"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 1 || digits == "" {
		return 0, errors.New("invalid input")
	}

	largestProduct := int64(math.Inf(-1))

	for i := 0; i+span <= len(digits); i++ {
		currSpan := digits[i : i+span]
		currProduct := int64(1)
		for _, r := range currSpan {
			currDigit, err := strconv.Atoi(string(r))
			if err != nil {
				return 0, errors.New("invalid input")
			}
			currProduct = currProduct * int64(currDigit)
		}
		if currProduct > largestProduct {
			largestProduct = currProduct
		}
	}

	return largestProduct, nil
}
