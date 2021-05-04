package core

var White [16]Piece
var Black [16]Piece

func InitializePlayers() {
	whiteRank := 0
	blackRank := 7
	// TODO: Better way of getting to each piece than random int
	White[0] = Piece{Name: "K", Square: Square{File: 4, Rank: whiteRank}, IsWhite: true}
	White[1] = Piece{Name: "Q", Square: Square{File: 3, Rank: whiteRank}, IsWhite: true}
	White[2] = Piece{Name: "B", Square: Square{File: 2, Rank: whiteRank}, IsWhite: true}
	White[3] = Piece{Name: "B", Square: Square{File: 5, Rank: whiteRank}, IsWhite: true}
	White[4] = Piece{Name: "N", Square: Square{File: 1, Rank: whiteRank}, IsWhite: true}
	White[5] = Piece{Name: "N", Square: Square{File: 6, Rank: whiteRank}, IsWhite: true}
	White[6] = Piece{Name: "R", Square: Square{File: 0, Rank: whiteRank}, IsWhite: true}
	White[7] = Piece{Name: "R", Square: Square{File: 7, Rank: whiteRank}, IsWhite: true}
	White[8] = Piece{Name: "P", Square: Square{File: 0, Rank: whiteRank + 1}, IsWhite: true}
	White[9] = Piece{Name: "P", Square: Square{File: 1, Rank: whiteRank + 1}, IsWhite: true}
	White[10] = Piece{Name: "P", Square: Square{File: 2, Rank: whiteRank + 1}, IsWhite: true}
	White[11] = Piece{Name: "P", Square: Square{File: 3, Rank: whiteRank + 1}, IsWhite: true}
	White[12] = Piece{Name: "P", Square: Square{File: 4, Rank: whiteRank + 1}, IsWhite: true}
	White[13] = Piece{Name: "P", Square: Square{File: 5, Rank: whiteRank + 1}, IsWhite: true}
	White[14] = Piece{Name: "P", Square: Square{File: 6, Rank: whiteRank + 1}, IsWhite: true}
	White[15] = Piece{Name: "P", Square: Square{File: 7, Rank: whiteRank + 1}, IsWhite: true}

	Black[0] = Piece{Name: "k", Square: Square{File: 4, Rank: blackRank}}
	Black[1] = Piece{Name: "q", Square: Square{File: 3, Rank: blackRank}}
	Black[2] = Piece{Name: "b", Square: Square{File: 2, Rank: blackRank}}
	Black[3] = Piece{Name: "b", Square: Square{File: 5, Rank: blackRank}}
	Black[4] = Piece{Name: "n", Square: Square{File: 1, Rank: blackRank}}
	Black[5] = Piece{Name: "n", Square: Square{File: 6, Rank: blackRank}}
	Black[6] = Piece{Name: "r", Square: Square{File: 0, Rank: blackRank}}
	Black[7] = Piece{Name: "r", Square: Square{File: 7, Rank: blackRank}}
	Black[8] = Piece{Name: "p", Square: Square{File: 0, Rank: blackRank - 1}}
	Black[9] = Piece{Name: "p", Square: Square{File: 1, Rank: blackRank - 1}}
	Black[10] = Piece{Name: "p", Square: Square{File: 2, Rank: blackRank - 1}}
	Black[11] = Piece{Name: "p", Square: Square{File: 3, Rank: blackRank - 1}}
	Black[12] = Piece{Name: "p", Square: Square{File: 4, Rank: blackRank - 1}}
	Black[13] = Piece{Name: "p", Square: Square{File: 5, Rank: blackRank - 1}}
	Black[14] = Piece{Name: "p", Square: Square{File: 6, Rank: blackRank - 1}}
	Black[15] = Piece{Name: "p", Square: Square{File: 7, Rank: blackRank - 1}}
}