package main

import (

	"../tic_tac_toe/src/game"
	"../tic_tac_toe/src/player"
)

func main() {
	player1 := player.Player{
		Name:   "ankon",
		Symbol: "X",
		Moves:  []int{},
	}
	player2 := player.Player{
		Name:   "chandan",
		Symbol: "O",
		Moves:  []int{},
	}
	game := game.Game{
		Players:             []player.Player{player1, player2},
		WinningCombinations: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9}, []int{1, 5, 9}, []int{3, 5, 7}},
	}
	game.MakeMove(1,0)
	game.MakeMove(4,1)
	game.MakeMove(2,0)
	game.MakeMove(5,1)
	game.MakeMove(7,0)
	game.MakeMove(6,1)
}
