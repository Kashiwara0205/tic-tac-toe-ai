package main

import (
	"code/game"
	"code/minmax_player"
	"code/random_player"
	"fmt"
)


type PlayerInterface interface{
	SelectPosition(g game.Game) (int, int)
}

type Players []PlayerInterface

func main() {
	game := game.CreateNewGame()
	minMaxPlayer := minmax_player.CreateMinMaxPlayer()
	randomPlayer := random_player.CreateRandomPlayer()

	var players Players
	players = append(players, &minMaxPlayer)
	players = append(players, &randomPlayer)

	for !game.CheckEnd(){
		row, col := players[game.GetPlayerTurn()].SelectPosition(game)
		game.PlaceTip(row, col)
		game.UpdateToNextTurn()
	}

	fmt.Printf("RESULT\n")
	fmt.Printf("----------------------------------\n")
	game.PrintBoard()
	fmt.Printf("minmax (x) is %v\n", formatMatchResult( game.GetMatchResult(0)) )
	fmt.Printf("random (o) is %v\n", formatMatchResult( game.GetMatchResult(1)) )
	fmt.Printf("----------------------------------\n")
}

func formatMatchResult(matchResult int) string{
	if 0 == matchResult { return "win" }
	if 1 == matchResult { return "lose" }
	if 2 == matchResult { return "draw" }

	return "none"
}