package core

import (
	"fmt"
)

func listLegalMoves(p Piece) []Square {
	legalMoves := []Square{}
	switch name := p.name; name {
	case 'p':
		var movement Square
		if p.square.rank == 1 && p.isWhite {
			movement = Square{p.square.file, p.square.rank + 2}
		} else if p.square.rank == 7 && !p.isWhite {
			movement = Square{p.square.file, p.square.rank - 2}
		}

		if board[movement.rank][fileMap[movement.file]] == 'x' && board[movement.rank-1][fileMap[movement.file]] == 'x' {
			legalMoves = append(legalMoves, movement)
		}

		if p.isWhite {
			movement = Square{p.square.file, p.square.rank + 1}
		} else {
			movement = Square{p.square.file, p.square.rank - 1}
		}

		if board[movement.rank][fileMap[movement.file]] == 'x' {
			legalMoves = append(legalMoves, movement)
		}
	}
	return legalMoves
}

func main() {
	initializePlayers()
	initializeBoard()
	printBoard()
	fmt.Println(listLegalMoves(white[14]))
}
