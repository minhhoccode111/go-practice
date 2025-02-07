package main

import (
	"fmt"
	"math"
)

/*
Range
The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The
first is the index, and the second is a copy of the element at that index.
*/

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
		fmt.Printf("2**%d = %f\n", i, math.Pow(2, float64(i)))
	}
}
