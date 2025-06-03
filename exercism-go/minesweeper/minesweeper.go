package minesweeper

import "strconv"

const mine byte = '*'

// Annotate returns an annotated board
func Annotate(board []string) []string {
	for i, v := range board {
		var line []byte

		for j, r := range v {
			if mine == byte(r) {
				line = append(line, mine)
				continue
			}

			var count byte = around(board, i, j)
			line = append(line, count)
		}
		board[i] = string(line)
	}

	return board
}

func around(board []string, i, j int) byte {
	count := 0
	items := [8][2]int{
		{i - 1, j - 1},
		{i - 1, j - 0},
		{i - 1, j + 1},
		{i - 0, j - 1},
		{i - 0, j + 1},
		{i + 1, j - 1},
		{i + 1, j - 0},
		{i + 1, j + 1},
	}

	for _, v := range items {
		if v[0] < 0 || v[1] < 0 || v[0] > len(board)-1 || v[1] > len(board[i])-1 {
			continue
		}

		if board[v[0]][v[1]] == byte(mine) {
			count++
		}
	}

	if count > 0 {
		return strconv.Itoa(count)[0]

	}

	return ' '
}
