package wordsearch

type Table map[string][2][2]int

type Direction int

const (
	N Direction = iota // 0
	NE
	E
	ES
	S
	SW
	W
	WN
)

func (d Direction) Offset(x, y int) (int, int) {
	switch d {
	case N:
		return x,
			y + 1
	case NE:
		return x + 1,
			y + 1
	case E:
		return x + 1,
			y
	case ES:
		return x + 1,
			y - 1
	case S:
		return x,
			y - 1
	case SW:
		return x - 1,
			y - 1
	case W:
		return x - 1,
			y
	case WN:
		return x - 1,
			y + 1
	default:
		return x,
			y
	}
}

func Solve(words []string, puzzle []string) (Table, error) {
	result := map[string][2][2]int{}
	return result, nil
}
