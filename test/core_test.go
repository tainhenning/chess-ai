package test

import (
	"chess-ai/core"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

type game struct {
	Board [8][8]string `json:"board"`
}

func TestPawnMovement(t *testing.T) {
	t.Run("Check Pawn movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.LoadBoard()
		var ans = core.White[12].LegalPawnMoves()
		var expected = []int {0, 1}
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check Pawn movement option 1", func(t *testing.T) {
		core.InitializePlayers()
		var expected = core.Square{File: 4, Rank: 2}
		core.White[12].MovePawn(1)
		var ans = core.White[12].Square
		if ans != expected {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check Pawn movement option 2", func(t *testing.T) {
		core.InitializePlayers()
		var expected = core.Square{File: 4, Rank: 3}
		core.White[12].MovePawn(2)
		var ans = core.White[12].Square
		if ans != expected {
			t.Errorf("got %d, want $%d", ans, expected)
			return
		}
	})
}

func TestBoardInitialization(t *testing.T) {
	t.Run("Check Board Setup", func(t *testing.T) {
		core.InitializePlayers()
		core.LoadBoard()
		// TODO: This is confusing because the board is currently being printed in black's perspective

		jsonFile, err := os.Open("./resources/pawn_movement_test.json")
		if err != nil {
			t.Errorf("Failed to open JSON file!")
			return
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var resource game
		if err2 := json.Unmarshal(byteValue, &resource); err2 != nil {
			t.Errorf("Unmarshal: %v\n", err)
		}
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				t.Errorf("Failed to close JSON")
				return
			}
		}(jsonFile)

		var ans = core.Board
		var expected = resource.Board
		if ans != expected {
			t.Errorf("got \n %s, want \n %s", ans, expected)
			return
		}
	})
}
