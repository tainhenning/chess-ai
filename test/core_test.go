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

func TestPieceMovement(t *testing.T) {
	// TODO: Compare individual moves
	t.Run("Check Pawn movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.InitializeBoard()
		ans := core.White[12].CheckLegalMoves()
		expected := []core.Square {{4, 2}, {4, 3}}
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check Knight movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.InitializeBoard()
		var ans = core.White[4].CheckLegalMoves()
		var expected = []core.Square {{2, 2}, {0, 2}}
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check Bishop movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.White[2].Square = core.Square{File: 3, Rank: 3}
		core.InitializeBoard()
		core.PrintBoard()
		var ans = core.White[2].CheckLegalMoves()
		var expected = []core.Square{
			{4, 4},
			{5, 5},
			{2, 2},
			{2, 4},
			{1, 5},
			{4, 2},
		}
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})

	t.Run("Check non pawn with pawn movement options", func(t *testing.T) {
		core.InitializePlayers()
		core.InitializeBoard()
		var ans = core.White[0].CheckLegalMoves()
		var expected []core.Square
		if len(ans) != len(expected) {
			t.Errorf("got %d, want %d", ans, expected)
			return
		}
	})
	//
	//t.Run("Check Pawn movement option 1", func(t *testing.T) {
	//	core.InitializePlayers()
	//	var expected = core.Square{File: 4, Rank: 2}
	//	core.White[12].MovePawn(1)
	//	var ans = core.White[12].Square
	//	if ans != expected {
	//		t.Errorf("got %d, want %d", ans, expected)
	//		return
	//	}
	//})
	//
	//t.Run("Check Pawn movement option 2", func(t *testing.T) {
	//	core.InitializePlayers()
	//	var expected = core.Square{File: 4, Rank: 3}
	//	core.White[12].MovePawn(2)
	//	var ans = core.White[12].Square
	//	if ans != expected {
	//		t.Errorf("got %d, want $%d", ans, expected)
	//		return
	//	}
	//})
}

func TestBoardInitialization(t *testing.T) {
	t.Run("Check Board Setup", func(t *testing.T) {
		core.InitializePlayers()
		core.InitializeBoard()
		// TODO: This is confusing because the board is currently being printed in black's perspective
		var ans = core.Board
		var expected = loadBoard("./resources/pawn_movement_test.json")
		if ans != expected {
			t.Errorf("got \n %s, want \n %s", ans, expected)
			return
		}
	})
}
