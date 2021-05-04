package core

import "fmt"

func PrintBoard() {
	for i := 7; i > -1; i-- {
		for j := 0; j < 8; j++ {
			fmt.Print(string(Board[i][j]), " ")
		}
		fmt.Print("\n")
	}
}
