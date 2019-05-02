package game

import (
	"fmt"
	"../player"
)

type Game struct {
	Players             []player.Player
	WinningCombinations [][]int
}

func (game Game) MakeMove(move int, playerId int) {
	game.Players[playerId].MakeMove(move)
	state := game.HasWon(playerId)
	if state {
		fmt.Println(game.Players[playerId].Name +" has won")
	}
}

func (game Game) HasWon(playerId int) bool {
	count := 0
	isMatching := false
	for count < len(game.WinningCombinations) {
		if game.Players[playerId].HasMoves(game.WinningCombinations[count]) {
			isMatching = true
			break
		}
		count++
	}
	return isMatching
}
