package core

// Piece represents the piece
type Piece struct {
	name    rune
	square  Square
	isWhite bool
}

func moveFile(p Piece, m int) {

}

func moveRank(p Piece, m int) {

}

func moveKnight(p Piece, option int) {
	switch o := option; o {
	case 1:
		moveFile(p, 1)
		moveRank(p, 2)
	case 2:
		moveFile(p, 1)
		moveRank(p, 2)
	case 4:
		moveFile(p, 2)
		moveRank(p, -1)
	case 5:
		moveFile(p, 1)
		moveRank(p, -2)
	case 7:
		moveFile(p, -1)
		moveRank(p, -2)
	case 8:
		moveFile(p, -2)
		moveRank(p, -1)
	case 10:
		moveFile(p, -2)
		moveRank(p, 1)
	case 11:
		moveFile(p, -1)
		moveRank(p, 2)
	}
}

func movePawn(p Piece, option int) {
	colorMod := 1
	if !p.isWhite {
		colorMod = -1
	}
	switch o := option; o {
	case 1:
		moveRank(p, colorMod)
	case 2:
		moveRank(p, colorMod*2)
	case 3:
		moveRank(p, colorMod)
		moveFile(p, -1)
	case 4:
		moveRank(p, colorMod)
		moveFile(p, 1)
	}
}
