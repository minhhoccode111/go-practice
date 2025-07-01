package dominoes

type Domino [2]int

// MakeChain tries to arrange dominoes in a circle with matching edges
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	n := len(input)
	if n == 0 {
		return []Domino{}, true
	}
	if n == 1 {
		if input[0][0] == input[0][1] {
			return input, true
		}
		return nil, false
	}

	// Helper: Try all permutations recursively
	used := make([]bool, n)
	res := make([]Domino, 0, n)
	var dfs func(pos int, end int) bool

	dfs = func(pos int, end int) bool {
		if pos == n {
			// Check circular: last connects to first
			return res[0][0] == end
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			// Try both orientations
			for flip := 0; flip < 2; flip++ {
				var left, right int
				if flip == 0 {
					left, right = input[i][0], input[i][1]
				} else {
					left, right = input[i][1], input[i][0]
				}
				if left == end {
					used[i] = true
					res = append(res, Domino{left, right})
					if dfs(pos+1, right) {
						return true
					}
					res = res[:len(res)-1]
					used[i] = false
				}
			}
		}
		return false
	}

	// Try all dominoes as starting point
	for i := 0; i < n; i++ {
		for flip := 0; flip < 2; flip++ {
			used[i] = true
			var left, right int
			if flip == 0 {
				left, right = input[i][0], input[i][1]
			} else {
				left, right = input[i][1], input[i][0]
			}
			res = append(res, Domino{left, right})
			if dfs(1, right) && res[0][0] == res[n-1][1] {
				// Chain found!
				// Make a copy to avoid side-effects
				out := make([]Domino, n)
				copy(out, res)
				return out, true
			}
			res = res[:0]
			used[i] = false
		}
	}
	return nil, false
}
