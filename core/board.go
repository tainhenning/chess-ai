package core

var Board [8][8]string

func LoadBoard() {

	for k := 0; k < 8; k++ {
		for l := 0; l < 8; l++ {
			Board[k][l] = "x"
		}
	}
	for i := 0; i < len(White); i++ {
		f := White[i].Square.File
		r := White[i].Square.Rank
		Board[r][f] = White[i].Name
	}

	for j := 0; j < len(Black); j++ {
		f := Black[j].Square.File
		r := Black[j].Square.Rank
		Board[r][f] = Black[j].Name
	}
}
