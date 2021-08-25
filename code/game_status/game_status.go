package game_status

import(
	"code/game"
)

// A class that conveys the game status to each player
type GameStatus struct{
	board [3][3] int
	playerTurn int
}

func CreateGameStatus(g game.Game) GameStatus{
	return GameStatus{ board: g.GetBoard(), playerTurn: g.GetCurrentPlayer() }
}

func (g *GameStatus) ExistsTip(row int, col int) bool{
	return game.NONE != g.board[row][col]
}

func (g *GameStatus) GetPlayerTurn() int{
	return g.playerTurn
}