package resistorcolortrio

import (
	"fmt"
	"math"
)

var dict = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	sum := 0
	for i, v := range colors {
		val, ok := dict[v]

		if !ok {
			panic("invalid color")
		}

		if i == 2 {
			sum *= int(math.Pow(10, float64(val)))
			continue
		}

		if i == 1 {
			sum += val
			continue
		}

		if i == 0 {
			sum += val * 10
			continue
		}
		// ignore extra colors
	}
	return format(sum)
}

func format(n int) string {
	switch {
	case n == 0:
		return fmt.Sprintf("0 ohms")
	case n%1_000_000_000 == 0:
		n /= 1_000_000_000
		return fmt.Sprintf("%d gigaohms", n)
	case n%1_000_000 == 0:
		n /= 1_000_000
		return fmt.Sprintf("%d megaohms", n)
	case n%1_000 == 0:
		n /= 1_000
		return fmt.Sprintf("%d kiloohms", n)
	default:
		return fmt.Sprintf("%d ohms", n)
	}
}
