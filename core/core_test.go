package core

import (
	"testing"
)

func TestPawnMovement(t *testing.T) {

}

func TestBoardInitialization(t *testing.T) {
	t.Run("Check Board Setup", func(t *testing.T) {
		initializePlayers()
		initializeBoard()
		var ans = board
		var expected = [8][8]rune{
			{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
			{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
			{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'},
			{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'},
			{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'},
			{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'},
			{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
			{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
		}
		if ans != expected {
			t.Errorf("got \n %d, want \n %d", ans, expected)
		}
	})
}

func TestLegalMoves(t *testing.T) {
	t.Run("Check Legal Moves for Pawn", func(t *testing.T) {
		initializePlayers()
		initializeBoard()
		var ans = listLegalMoves(white[14])
		if len(ans) < 2 {
			t.Errorf("got %d, want %d", ans, 2)
		}
	})
}
