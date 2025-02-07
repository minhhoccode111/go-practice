package main

import (
	"fmt"
	"strings"
)

/*
Slices of slices
Slices can contain any type, including other slices.
*/

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"-", "-", "-", "-", "-"},
		{"|", " ", " ", " ", "|"},
		{"|", " ", " ", " ", "|"},
		{"|", " ", " ", " ", "|"},
		{"-", "-", "-", "-", "-"},
	}

	// The players take turns.
	board[1][1] = "X"
	board[3][3] = "O"
	board[2][3] = "X"
	board[2][1] = "O"
	board[1][3] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
