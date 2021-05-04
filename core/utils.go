package core

import "fmt"

var fileMap = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
}

func printBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Print("\n")
	}
}
