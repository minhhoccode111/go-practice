package railfence

func Encode(message string, rails int) string {
	array := make([][]rune, rails)

	for i := range array {
		array[i] = make([]rune, len(message))
	}

	row := 0
	isDown := true
	for i, r := range message {
		array[row][i] = r
		if isDown {
			row++
			if row == len(array) {
				row -= 2
				isDown = false
			}
			continue
		} else {
			row--
			if row == -1 {
				row += 2
				isDown = true
			}
		}
	}

	result := ""

	for _, row := range array {
		var tmp []rune
		for _, r := range row {
			if r != 0 {
				tmp = append(tmp, r)
			}
		}
		result += string(tmp)
	}

	return result
}

func Decode(message string, rails int) string {
	mark := make([][]bool, rails)
	for i := range mark {
		mark[i] = make([]bool, len(message))
	}

	row, down := 0, true
	for col := range message {
		mark[row][col] = true
		if down {
			row++
			if row == rails {
				row = rails - 2
				down = false
			}
		} else {
			row--
			if row < 0 {
				row = 1
				down = true
			}
		}
	}

	res := make([][]rune, rails)
	for i := range res {
		res[i] = make([]rune, len(message))
	}

	idx := 0
	for i := 0; i < rails; i++ {
		for j := 0; j < len(message); j++ {
			if mark[i][j] {
				res[i][j] = rune(message[idx])
				idx++
			}
		}
	}

	row, down = 0, true
	var decoded []rune
	for col := 0; col < len(message); col++ {
		decoded = append(decoded, res[row][col])
		if down {
			row++
			if row == rails {
				row = rails - 2
				down = false
			}
		} else {
			row--
			if row < 0 {
				row = 1
				down = true
			}
		}
	}

	return string(decoded)
}

/*
row = 4
w.....w.....w.....w
.w...w.w...w.w...w.
..w.w...w.w...w.w..
...w.....w.....w...
*/
