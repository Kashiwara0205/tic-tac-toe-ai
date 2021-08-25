package minmax_player

import(
	"code/game"
)

type MinMaxPlayer struct{ }

func CreateMinMaxPlayer() MinMaxPlayer{
	return MinMaxPlayer{}
}

/*
func minxmax(g game.Game, depth int, myTurn int){
	if g.CheckEnd(){ return evaluate(g, depth, myTurn) }
	depth += 1

	scores := []

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if ( NONE == g.GetTipStatus(i, j) ){
				possibleGame := game.CopyGame()
				possibleGame.PlaceTip(i, j)
				possibleGame.UpdateToNextTurn()
				append(scores, minxmax(possibleGame, depth, myTurn) )
			}
		}
	}

	if ( myTurn == g.GetPlayerTurn() ){
		maxScoreIdx := 0
		return scores[maxScoreIdx]
	}else{
		minScoreIdx := 0
		return scores[minScoreIdx]
	}
}

*/
func (player *MinMaxPlayer) evaluate(g game.Game, depth int, myTurn int) int {
	if myTurn == g.GetPlayerTurn() && g.CheckWin() { return 10 - depth }
	if myTurn == g.GetPlayerTurn() && g.CheckLose() { return depth -10 }

	if myTurn != g.GetPlayerTurn() && g.CheckWin() { return depth -10 }
	if myTurn != g.GetPlayerTurn() && g.CheckLose() { return 10 - depth }

	return 0
}
