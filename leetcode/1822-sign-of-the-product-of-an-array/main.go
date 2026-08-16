package main

func arraySign(nums []int) int {
	count := 0
	for _, n := range nums {
		if n == 0 {
			return 0
		}

		if n < 0 {
			count++
		}
	}

	if count%2 == 0 {
		return 1
	}

	return -1
}
