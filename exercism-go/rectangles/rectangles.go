package rectangles

func Count(diagram []string) int {
	grid := diagram
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	count := 0
	// loop through rows with y, y1 means top-horizontal edge, y2 means bottom-horizontal edge
	for y1 := 0; y1 < rows; y1++ {
		// loop through cols with x, x1 means left-vertical edge, x2 means right-vertical edge
		for x1 := 0; x1 < cols; x1++ {
			// continue if not a '+' char (identify the start of a rectangle)
			if grid[y1][x1] != '+' {
				continue
			}

			// curr char is '+'

			// next char right of x1 (horizontal)
			for x2 := x1 + 1; x2 < cols; x2++ {
				// continue if not a '+' char or x1 and x2 not horizontal edge
				if grid[y1][x2] != '+' || !isHorizontalEdge(grid[y1], x1, x2) {
					continue
				}

				// x2 is '+' or horizontal edge with x1

				// next char down of y1 (vertical)
				for y2 := y1 + 1; y2 < rows; y2++ {
					// continue if x1 or x2 on this row is not '+'
					if grid[y2][x1] != '+' || grid[y2][x2] != '+' {
						continue
					}

					// x1 and x2 on this row is '+'

					// check if we can form a vertical edge between y1 and y2 at x1 column
					if isVerticalEdge(grid, x1, y1, y2) &&
						// check if we can form a vertical edge between y1 and y2 at x2 column
						isVerticalEdge(grid, x2, y1, y2) &&
						// check if we can form a horizontal edge between x1 and x2 at y1 row
						isHorizontalEdge(grid[y2], x1, x2) {
						count++
					}
				}
			}
		}
	}
	return count
}

// isHorizontalEdge take in a string and two indexes to check if characters
// between two indexes of a string are valid (inclusive '-' and '+')
func isHorizontalEdge(row string, x1, x2 int) bool {
	for x := x1 + 1; x < x2; x++ {
		if row[x] != '-' && row[x] != '+' {
			return false
		}
	}
	return true
}

// isVerticalEdge take in a slice of strings and three indexes to check if
// characters between the two indexes of specified column are valid (inclusive
// '|' and '+')
func isVerticalEdge(grid []string, col, y1, y2 int) bool {
	for y := y1 + 1; y < y2; y++ {
		c := grid[y][col]
		if c != '|' && c != '+' {
			return false
		}
	}
	return true
}
