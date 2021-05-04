package core

var board [8][8]rune

func initializeBoard() {

	for k := 0; k < 8; k++ {
		for l := 0; l < 8; l++ {
			board[k][l] = 'x'
		}
	}

	for i := 0; i < len(white); i++ {
		f := fileMap[white[i].square.file]
		r := white[i].square.rank
		board[r][f] = white[i].name
	}

	for j := 0; j < len(black); j++ {
		f := fileMap[black[j].square.file]
		r := black[j].square.rank
		board[r][f] = black[j].name
	}
}

// Board gets the current value of the board
func Board() [8][8]rune {
	return board
}

// SetBoard sets the value of board
func SetBoard(b [8][8]rune) {
	board = b
}
