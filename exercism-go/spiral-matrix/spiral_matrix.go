package spiralmatrix

type Direction string

const (
	Right Direction = "right"
	Left  Direction = "left"
	Down  Direction = "down"
	Up    Direction = "up"
)

func (s Direction) Next(currI, currJ int) (int, int, Direction) {
	switch s {
	case Right:
		return currI + 1, currJ, Down
	case Down:
		return currI, currJ - 1, Left
	case Left:
		return currI - 1, currJ, Up
	case Up:
		return currI, currJ + 1, Right
	default:
		panic("invalid direction")
	}
}

func SpiralMatrix(size int) [][]int {
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}
	going(size, 1, 0, 0, Right, result)
	return result
}

func going(size, curr, i, j int, d Direction, matrix [][]int) {
	if i >= size || j >= size || i < 0 || j < 0 || matrix[i][j] != 0 {
		return
	}
	switch d {
	case Right:
		for j < size {
			if matrix[i][j] != 0 {
				break
			}
			matrix[i][j] = curr
			curr++
			j++
		}
		j--
	case Down:
		for i < size {
			if matrix[i][j] != 0 {
				break
			}
			matrix[i][j] = curr
			curr++
			i++
		}
		i--
	case Left:
		for j >= 0 {
			if matrix[i][j] != 0 {
				break
			}
			matrix[i][j] = curr
			curr++
			j--
		}
		j++
	case Up:
		for i >= 0 {
			if matrix[i][j] != 0 {
				break
			}
			matrix[i][j] = curr
			curr++
			i--
		}
		i++
	}
	nextI, nextJ, nextD := d.Next(i, j)
	going(size, curr, nextI, nextJ, nextD, matrix)
}
