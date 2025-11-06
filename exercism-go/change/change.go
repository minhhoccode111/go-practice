package change

import (
	"errors"
)

func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, errors.New("taivisao")
	}

	dp := make([][]int, target+1)
	dp[0] = []int{}
	for i := 1; i <= target; i++ {
		var best []int
		for _, c := range coins {
			if i >= c && dp[i-c] != nil {
				current := append([]int{}, dp[i-c]...)
				current = append([]int{c}, current...)
				if best == nil || len(current) < len(best) {
					best = current
				}
			}
		}
		dp[i] = best
	}

	if dp[target] == nil {
		return nil, errors.New("taivisao")
	}

	return dp[target], nil
}

/*

coins = [1, 4, 5]
target = 8

  1 2   3 4 5 6 7 8 -
1 1 1 1             -
4                   -
5                   -

*/
