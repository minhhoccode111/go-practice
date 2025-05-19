package yacht

func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return anyCombination(1, dice)
	case "twos":
		return anyCombination(2, dice)
	case "threes":
		return anyCombination(3, dice)
	case "fours":
		return anyCombination(4, dice)
	case "fives":
		return anyCombination(5, dice)
	case "sixes":
		return anyCombination(6, dice)
	case "full house":
		if isFullHouse(dice) {
			return sumOfDice(dice)
		}
		return 0
	case "four of a kind":
		if count, value := countMostFrequent(dice); count >= 4 {
			return value * 4
		}
		return 0
	case "little straight":
		if isLittleStraight(dice) {
			return 30
		}
		return 0
	case "big straight":
		if isBigStraight(dice) {
			return 30
		}
		return 0
	case "choice":
		return sumOfDice(dice)
	case "yacht":
		if isYacht(dice) {
			return 50
		}
		return 0
	default:
		return 0
	}
}

func isYacht(dice []int) bool {
	return dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4]
}

func isFullHouse(dice []int) bool {
	counts := make(map[int]int)
	for _, d := range dice {
		counts[d]++
	}
	if len(counts) != 2 {
		return false
	}
	for _, count := range counts {
		if count != 2 && count != 3 {
			return false
		}
	}
	return true
}

func countMostFrequent(dice []int) (int, int) {
	counts := make(map[int]int)
	maxCount := 0
	maxValue := 0
	for _, d := range dice {
		counts[d]++
		if counts[d] > maxCount {
			maxCount = counts[d]
			maxValue = d
		}
	}
	return maxCount, maxValue
}

func isLittleStraight(dice []int) bool {
	seen := make(map[int]bool)
	for _, d := range dice {
		if d < 1 || d > 5 {
			return false
		}
		seen[d] = true
	}
	return len(seen) == 5
}

func isBigStraight(dice []int) bool {
	seen := make(map[int]bool)
	for _, d := range dice {
		if d < 2 || d > 6 {
			return false
		}
		seen[d] = true
	}
	return len(seen) == 5
}

func sumOfDice(dice []int) int {
	sum := 0
	for _, v := range dice {
		sum += v
	}
	return sum
}

func anyCombination(n int, dice []int) int {
	sum := 0
	for _, v := range dice {
		if v == n {
			sum += n
		}
	}
	return sum
}
