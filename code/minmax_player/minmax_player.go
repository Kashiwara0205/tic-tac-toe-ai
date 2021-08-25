package minmax_player

import(
	"code/game"
)

type MinMaxPlayer struct{ }

func CreateMinMaxPlayer() MinMaxPlayer{
	return MinMaxPlayer{}
}

func (player *MinMaxPlayer) evaluate(g game.Game) int {
	if g.CheckWin() { return 10 }
	if g.CheckLose() { return -10 }

	return 0
}
