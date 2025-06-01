package stateoftictactoe

import (
	"errors"
	"fmt"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	dict := map[int]rune{}
	count := 1

	for _, v := range board {
		for _, r := range v {
			dict[count] = r
			count++

			if r == 'X' {
				dict['X']++
			}

			if r == 'O' {
				dict['O']++
			}

			if r == ' ' {
				dict[' ']++
			}
		}
	}

	if dict['X']-dict['O'] < 0 || dict['X']-dict['O'] > 1 {
		return "", errors.New("given board cannot be reached")
	}

	if isWin(dict, 'X') && isWin(dict, 'O') {
		return "", errors.New("the game was plated after it already ended")
	}

	if dict[' '] == 0 && !isWin(dict, 'X') && !isWin(dict, 'O') {
		return Draw, nil
	}

	if !isWin(dict, 'X') && !isWin(dict, 'O') {
		return Ongoing, nil
	}

	fmt.Println(dict)

	return Win, nil
}

func isWin(dict map[int]rune, r rune) bool {
	rowOne := (r == dict[1] && r == dict[2] && r == dict[3])
	rowTwo := (r == dict[4] && r == dict[5] && r == dict[6])
	rowThree := (r == dict[7] && r == dict[8] && r == dict[9])
	columnOne := (r == dict[1] && r == dict[4] && r == dict[7])
	columnTwo := (r == dict[2] && r == dict[5] && r == dict[8])
	columnThree := (r == dict[3] && r == dict[6] && r == dict[9])
	diagonalOne := (r == dict[1] && r == dict[5] && r == dict[9])
	diagonalTwo := (r == dict[3] && r == dict[5] && r == dict[7])

	return columnOne || rowOne || columnTwo || rowTwo || columnThree || rowThree || diagonalOne || diagonalTwo
}
