package test

import (
	"chess-ai/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type game struct {
	Board [8][8]string `json:"board"`
}

func loadBoard(filePath string) [8][8]string {
	var resource game
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open JSON file")
		return resource.Board
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err2 := json.Unmarshal(byteValue, &resource); err2 != nil {
		fmt.Printf("Unmarshal: %v\n", err)
		return resource.Board
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("Failed to close JSON")
		}
	}(jsonFile)
	return resource.Board
}

func TestPawnMovement(t *testing.T) {
	// TODO: Compare individual moves
	t.Run("Check Pawn movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.LoadBoard()
		var ans = core.White[12].CheckLegalMoves()
		var expected = []int {0, 1}
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check Knight movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.LoadBoard()
		var ans = core.White[4].CheckLegalMoves()
		var expected = []int {0, 1}
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check non pawn with pawn movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.LoadBoard()
		var ans = core.White[0].CheckLegalMoves()
		var expected []int
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
		var ans = core.Board
		var expected = loadBoard("./resources/pawn_movement_test.json")
		if ans != expected {
			t.Errorf("got \n %s, want \n %s", ans, expected)
			return
		}
	})
}
