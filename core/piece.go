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

func (p *Piece) CheckLegalMoves() []int {
	// TODO: Potential bottleneck with ToLower
	// TODO: Translate legal moves from options to square
	// TODO: Figure out how to merge legal*Moves functions
	switch name := strings.ToLower(p.Name); name {
	case "p":
		return p.legalPawnMoves(pawnMoves)
	case "n":
		return p.legalKnightMoves(knightMoves)
	}
	return []int {}
}

func (p *Piece) legalPawnMoves(moveLimit int) []int {
	var legalMoves []int
	for i := 0; i < moveLimit; i++ {
		if p.checkPawnMoveLegality(i) {
			legalMoves = append(legalMoves, i)
		}
	}
	return legalMoves
}

func (p *Piece) legalKnightMoves(moveLimit int) []int {
	var legalMoves []int
	for i := 0; i < moveLimit; i++ {
		if p.checkKnightMoveLegality(i) {
			legalMoves = append(legalMoves, i)
		}
	}
	return legalMoves
}

func (p *Piece) examineSquareValidity (rank int, file int) bool {
	newRank := p.Square.Rank + rank
	newFile := p.Square.File + file
	// TODO: Will need to change here for captures
	if newRank > -1 && newRank < 8 && newFile > -1 && newFile < 8 && Board[newRank][newFile] == "x" {
		return true
	}
	return false
}

func (p *Piece) checkKnightMoveLegality(option int) bool {
	switch o := option; o {
	case 0:
		return p.examineSquareValidity(1, 2)
	case 1:
		return p.examineSquareValidity(2, 1)
	case 2:
		return p.examineSquareValidity(2, -1)
	case 3:
		return p.examineSquareValidity(1, -2)
	case 4:
		return p.examineSquareValidity(-1, -2)
	case 5:
		return p.examineSquareValidity(-2, -1)
	case 6:
		return p.examineSquareValidity(-2, 1)
	case 7:
		return p.examineSquareValidity(-1, 2)
	}
	return false
}

func (p *Piece) moveKnight(option int) {
	switch o := option; o {
	case 0:
		p.moveFile(1)
		p.moveRank(2)
	case 1:
		p.moveFile(2)
		p.moveRank(1)
	case 2:
		p.moveFile(2)
		p.moveRank(-1)
	case 3:
		p.moveFile(1)
		p.moveRank(-2)
	case 4:
		p.moveFile(-1)
		p.moveRank(-2)
	case 5:
		p.moveFile(-2)
		p.moveRank(-1)
	case 6:
		p.moveFile(-2)
		p.moveRank(1)
	case 7:
		p.moveFile(-1)
		p.moveRank(2)
	}
}

// TODO: Merge logic of checkPawnMoveLegality and MovePawn
func (p *Piece) checkPawnMoveLegality(option int) bool {
	colorMod := 1
	if !p.IsWhite {
		colorMod = -1
	}
	switch o := option; o {
	case 0:
		if Board[p.Square.Rank+colorMod][p.Square.File] == "x" {
			return true
		}
	case 1:
		if Board[p.Square.Rank + (2*colorMod)][p.Square.File] == "x" {
			return true
		}
	case 2:
		if Board[p.Square.Rank + colorMod][p.Square.File - 1] != "x" {
			return true
		}
	case 3:
		if Board[p.Square.Rank + colorMod][p.Square.File + 1] != "x" {
			return true
		}
	}
	return false
}

func (p *Piece) MovePawn(option int) {
	colorMod := 1
	if !p.IsWhite {
		colorMod = -1
	}
	switch o := option; o {
	case 1:
		p.moveRank(colorMod)
	case 2:
		p.moveRank(colorMod*2)
	case 3:
		p.moveRank(colorMod)
		p.moveFile(-1)
	case 4:
		p.moveRank(colorMod)
		p.moveFile(1)
	}
}
