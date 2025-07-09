package twobucket

import (
	"errors"
	// "fmt"
)

type state struct {
	bucket1 int
	bucket2 int
}

// Solve determines the number of moves to reach the goal amount using two
// buckets
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (goalBucket string, moves int, otherBucket int, err error) {
	// edge cases
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 {
		return "", 0, 0, errors.New("invalid input: bucket sizes and goal must be > 0")
	}

	switch startBucket {
	case "one":
		return simulate(sizeBucketOne, sizeBucketTwo, goalAmount, "one")
	case "two":
		return simulate(sizeBucketTwo, sizeBucketOne, goalAmount, "two")
	default:
		return "", 0, 0, errors.New("startBucket must be 'one' or 'two'")
	}
}

// Simulate tries all possible actions you can take with the two buckets to
// reach the goal amount step by step, and returns the result when successful
// It uses a BFS strategy to explore all possible valid states (combinations of
// water in the two buckets) and tracks how many actions it takes
func simulate(startSize, otherSize, goal int, startName string) (string, int, int, error) {
	visited := make(map[state]bool)
	type step struct {
		b1, b2 int
		count  int
	}
	queue := []step{{startSize, 0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		s, o, cnt := current.b1, current.b2, current.count

		if visited[state{s, o}] {
			continue
		}
		visited[state{s, o}] = true

		// Goal check
		if s == goal {
			return startName, cnt, o, nil
		}
		if o == goal {
			goalName := "two"
			if startName == "two" {
				goalName = "one"
			}
			return goalName, cnt, s, nil
		}

		// Invalid state (rule: can't have start empty and other full)
		if s == 0 && o == otherSize {
			continue
		}

		// All valid next states (6 total)
		queue = append(queue,
			step{startSize, o, cnt + 1}, // Fill start
			step{0, o, cnt + 1},         // Empty start
			step{s, otherSize, cnt + 1}, // Fill other
			step{s, 0, cnt + 1},         // Empty other,
		)

		// Pour start → other
		pourToOther := min(s, otherSize-o)
		queue = append(queue, step{s - pourToOther, o + pourToOther, cnt + 1})

		// Pour other → start
		pourToStart := min(o, startSize-s)
		queue = append(queue, step{s + pourToStart, o - pourToStart, cnt + 1})
	}

	return "", 0, 0, errors.New("no solution found")
}

// Min compares between two integers and return the smaller one
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
