package core

import "strings"

// Piece represents the piece
type Piece struct {
	Name    string
	Square  Square
	IsWhite bool
}

var pawnMoves = 2
var knightMoves = 8

func (p *Piece) moveFile(m int) {
	p.Square.File += m
}

func (p *Piece) moveRank(m int) {
	p.Square.Rank += m
}

func (p *Piece) CheckLegalMoves() []Square {
	// TODO: Potential bottleneck with ToLower
	switch name := strings.ToLower(p.Name); name {
	case "p":
		return p.legalPawnMoves(pawnMoves)
	case "n":
		return p.legalKnightMoves(knightMoves)
	case "b":
		return p.legalBishopMoves()
	}
	return []Square {}
}

func (p *Piece) legalPawnMoves(moveLimit int) []Square {
	var legalMoves []Square
	var nullMove = Square{}

	for i := 0; i < moveLimit; i++ {
		move := p.checkPawnMoveLegality(i)
		if move != nullMove {
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

func (p *Piece) legalKnightMoves(moveLimit int) []Square {
	var legalMoves []Square
	var nullMove = Square{}
	for i := 0; i < moveLimit; i++ {
		move := p.checkKnightMoveLegality(i)
		if move != nullMove {
			legalMoves = append(legalMoves, move)
		}
	}
	return legalMoves
}

// TODO: Change input to a single square struct
// TODO: Also it is confusing that the inputs are being added
func (p *Piece) examineSquareValidity (file int, rank int) bool {
	newRank := p.Square.Rank + rank
	newFile := p.Square.File + file
	// TODO: Will need to change here for captures
	if newRank > -1 && newRank < 8 && newFile > -1 && newFile < 8 && Board[newRank][newFile] == "x" {
		return true
	}
	return false
}

func (p *Piece) legalBishopMoves() []Square {
	// TODO: I really don't like how this is setup at all
	var legalMoves []Square
	for i := 1; i + p.Square.File < 8 && i + p.Square.Rank < 8; i++ {
		if p.examineSquareValidity(i, i) {
			legalMoves = append(legalMoves, Square{p.Square.File + i, p.Square.Rank + i})
		} else {
			break
		}
	}
	for i := -1; i + p.Square.File > -1 && i + p.Square.Rank > -1; i-- {
		if p.examineSquareValidity(i, i) {
			legalMoves = append(legalMoves, Square{p.Square.File + i, p.Square.Rank + i})
		} else {
			break
		}
	}
	for i := -1; i + p.Square.File > -1 && i + p.Square.Rank > -1 && i + p.Square.File < 8 && i + p.Square.Rank < 8; i-- {
		if p.examineSquareValidity(i, -i) {
			legalMoves = append(legalMoves, Square{p.Square.File + i, p.Square.Rank - i})
		} else {
			break
		}
	}
	for i := -1; i + p.Square.File > -1 && i + p.Square.Rank > -1 && i + p.Square.File < 8 && i + p.Square.Rank < 8; i-- {
		if p.examineSquareValidity(-i, i) {
			legalMoves = append(legalMoves, Square{p.Square.File - i, p.Square.Rank + i})
		} else {
			break
		}
	}

	return legalMoves
}

func (p *Piece) checkKnightMoveLegality(option int) Square {
	switch o := option; o {
	case 0:
		if p.examineSquareValidity(2, 1) {
			return Square{p.Square.File + 2, p.Square.Rank + 1}
		}
	case 1:
		if p.examineSquareValidity(1, 2) {
			return Square{p.Square.File + 1, p.Square.Rank + 2}
		}
	case 2:
		if p.examineSquareValidity(-1, 2) {
			return Square{p.Square.File - 1, p.Square.Rank + 2}
		}
	case 3:
		if p.examineSquareValidity(-2, 1) {
			return Square{p.Square.File - 2, p.Square.Rank + 1}
		}
	case 4:
		if p.examineSquareValidity(-2, -1) {
			return Square{p.Square.File - 2, p.Square.Rank - 1}
		}
	case 5:
		if p.examineSquareValidity(-1, -2) {
			return Square{p.Square.File - 1, p.Square.Rank - 2}
		}
	case 6:
		if p.examineSquareValidity(1, -2) {
			return Square{p.Square.File + 1, p.Square.Rank - 2}
		}
	case 7:
		if p.examineSquareValidity(2, -1) {
			return Square{p.Square.File + 2, p.Square.Rank - 1}
		}
	}
	return Square{}
}

func (p *Piece) checkPawnMoveLegality(option int) Square {
	colorMod := 1
	if !p.IsWhite {
		colorMod = -1
	}
	switch o := option; o {
	case 0:
		if Board[p.Square.Rank + colorMod][p.Square.File] == "x" {
			return Square{p.Square.File, p.Square.Rank + colorMod}
		}
	case 1:
		if Board[p.Square.Rank + (2*colorMod)][p.Square.File] == "x" {
			return Square{p.Square.File, p.Square.Rank + (2 * colorMod)}
		}
	case 2:
		if Board[p.Square.Rank + colorMod][p.Square.File - 1] != "x" {
			return Square{p.Square.File - 1, p.Square.Rank + colorMod}
		}
	case 3:
		if Board[p.Square.Rank + colorMod][p.Square.File + 1] != "x" {
			return Square{p.Square.File + 1, p.Square.Rank + colorMod}
		}
	}
	return Square{}
}
