package core

var white [16]Piece
var black [16]Piece

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

// White gets white pieces
func White() [16]Piece {
	return white
}

// SetWhite sets white pieces
func SetWhite(w [16]Piece) {
	white = w
}

// Black gets black pieces
func Black() [16]Piece {
	return black
}

// SetBlack sets black pieces
func SetBlack(b [16]Piece) {
	black = b
}
