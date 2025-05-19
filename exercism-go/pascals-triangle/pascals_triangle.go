package pascal

/*

n = 4
[1]
[1 1]
[1 2 1]
[1 3 3 1]
*/

func Triangle(n int) [][]int {
	// init slices
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, i+1)
	}

	// inject values
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			tmp := 1
			if j != i && j != 0 {
				tmp = result[i-1][j] + result[i-1][j-1]
			}
			result[i][j] = tmp
		}
	}

	return result
}
