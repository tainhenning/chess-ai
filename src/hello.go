package main

import (
	"fmt"
)

// Square represents a position on the board
type Square struct {
	file rune
	rank int
}

// Piece represents the piece
type Piece struct {
	name    rune
	square  Square
	isWhite bool
}

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

var board [8][8]rune

var white [16]Piece
var black [16]Piece

func appendLegalMove(moves []Square, movement Square) []Square {
	if board[movement.rank][fileMap[movement.file]] == 'x' {
		moves = append(moves, movement)
		return moves
	}
	return moves
}

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

func printBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Print("\n")
	}
}

func initializePlayers() {
	whiteRank := 0
	blackRank := 7

	white[0] = Piece{'K', Square{'d', whiteRank}, true}
	white[1] = Piece{'Q', Square{'e', whiteRank}, true}
	white[2] = Piece{'B', Square{'c', whiteRank}, true}
	white[3] = Piece{'B', Square{'f', whiteRank}, true}
	white[4] = Piece{'N', Square{'b', whiteRank}, true}
	white[5] = Piece{'N', Square{'g', whiteRank}, true}
	white[6] = Piece{'R', Square{'a', whiteRank}, true}
	white[7] = Piece{'R', Square{'h', whiteRank}, true}
	white[8] = Piece{'P', Square{'d', whiteRank + 1}, true}
	white[9] = Piece{'P', Square{'e', whiteRank + 1}, true}
	white[10] = Piece{'P', Square{'c', whiteRank + 1}, true}
	white[11] = Piece{'P', Square{'f', whiteRank + 1}, true}
	white[12] = Piece{'P', Square{'b', whiteRank + 1}, true}
	white[13] = Piece{'P', Square{'g', whiteRank + 1}, true}
	white[14] = Piece{'P', Square{'a', whiteRank + 1}, true}
	white[15] = Piece{'P', Square{'h', whiteRank + 1}, true}

	black[0] = Piece{'K', Square{'d', blackRank}, false}
	black[1] = Piece{'Q', Square{'e', blackRank}, false}
	black[2] = Piece{'B', Square{'c', blackRank}, false}
	black[3] = Piece{'B', Square{'f', blackRank}, false}
	black[4] = Piece{'N', Square{'b', blackRank}, false}
	black[5] = Piece{'N', Square{'g', blackRank}, false}
	black[6] = Piece{'R', Square{'a', blackRank}, false}
	black[7] = Piece{'R', Square{'h', blackRank}, false}
	black[8] = Piece{'P', Square{'d', blackRank - 1}, false}
	black[9] = Piece{'P', Square{'e', blackRank - 1}, false}
	black[10] = Piece{'P', Square{'c', blackRank - 1}, false}
	black[11] = Piece{'P', Square{'f', blackRank - 1}, false}
	black[12] = Piece{'P', Square{'b', blackRank - 1}, false}
	black[13] = Piece{'P', Square{'g', blackRank - 1}, false}
	black[14] = Piece{'P', Square{'a', blackRank - 1}, false}
	black[15] = Piece{'P', Square{'h', blackRank - 1}, false}
}

func main() {
	initializePlayers()
	initializeBoard()
	printBoard()
	fmt.Println(listLegalMoves(white[14]))
}
