package diamond

import (
	"errors"
	"strings"
)

/*

Diamond for letter 'E': (n = 9)

····A···· 0 0
···B·B··· 1 1
··C···C·· 2 2
·D·····D· 3 3
E·······E 4 4 <- mid
·D·····D· 3 5
··C···C·· 2 6
···B·B··· 1 7
····A···· 0 8

*/

func Gen(char byte) (string, error) {
	mid := int(char - 'A')
	if mid > 25 || mid < 0 {
		return "", errors.New("invalid input")
	}
	n := mid*2 + 1
	t := createTable(n)
	for i := 0; i < n; i++ {
		var currChar string
		var space int
		if i <= mid {
			space = i
			currChar = string(rune(i + 'A'))
		} else {
			space = n - i - 1
			currChar = string(rune(space + 'A'))
		}

		for j := 0; j < n; j++ {
			forward := mid + space
			backward := mid - space
			if j == forward || j == backward {
				t[i][j] = currChar
			}
		}
	}
	return createResultFromTable(t), nil
}

func createResultFromTable(table [][]string) string {
	var resultSlice []string
	for _, s := range table {
		line := ""
		for _, v := range s {
			line += v
		}
		resultSlice = append(resultSlice, line)
	}
	return strings.Join(resultSlice, "\n")
}

func createTable(n int) [][]string {
	table := make([][]string, n)
	for i := 0; i < n; i++ {
		table[i] = make([]string, n)
		for j := 0; j < n; j++ {
			table[i][j] = " "
		}
	}
	return table
}
