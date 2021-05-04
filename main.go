package main

import (
	"chess-ai/core"
	"fmt"
)

func main() {
	core.InitializePlayers()
	core.LoadBoard()
	core.PrintBoard()
	fmt.Println("--------")
}
