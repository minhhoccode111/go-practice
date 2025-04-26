package darts

import "math"

/*
- If the dart lands outside the target, player earns no points (0 points).
- If the dart lands in the outer circle of the target, player earns 1 point.
- If the dart lands in the middle circle of the target, player earns 5 points.
- If the dart lands in the inner circle of the target, player earns 10 points.
*/

func Score(x, y float64) int {
	distanceToCenter := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case distanceToCenter <= 1:
		return 10
	case distanceToCenter <= 5:
		return 5
	case distanceToCenter <= 10:
		return 1
	default:
		return 0
	}
}
