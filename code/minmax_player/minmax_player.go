package minmax_player

import(
	"code/game"
	"code/utils"
)

type MinMaxPlayer struct{
	bestPosition [2]int
}

func CreateMinMaxPlayer() MinMaxPlayer{
	return MinMaxPlayer{}
}

func (player *MinMaxPlayer) SelectPosition(g game.Game) (int, int) {
	player.minmax(g, 0, g.GetPlayerTurn())

	row := player.bestPosition[0]
	col := player.bestPosition[1]

	return row, col
}

func (player *MinMaxPlayer) minmax(g game.Game, depth int, myTurn int) int {
	if g.CheckEnd(){ return evaluate(g, depth, myTurn) }
	depth += 1

	var scores[]int
	var positions[][]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if ( g.CanPlaceTip(i, j) ){
				possibleGame := g.CopyGame()
				possibleGame.PlaceTip(i, j)
				possibleGame.UpdateToNextTurn()
				scores = append(scores, player.minmax(possibleGame, depth, myTurn) )
				position := []int{i, j}
				positions = append(positions, position)
			}
		}
	}

	if ( myTurn == g.GetPlayerTurn() ){
		maxScoreIdx := utils.GetMaxIdx(scores)
		player.bestPosition[0] = positions[maxScoreIdx][0]
		player.bestPosition[1] = positions[maxScoreIdx][1]

		return scores[maxScoreIdx]
	}else{
		minScoreIdx := utils.GetMinIdx(scores)
		player.bestPosition[0] = positions[minScoreIdx][0]
		player.bestPosition[1] = positions[minScoreIdx][1]
		
		return scores[minScoreIdx]
	}
}

func evaluate(g game.Game, depth int, myTurn int) int {
	if myTurn == g.GetPlayerTurn() && g.CheckWin() { return 10 - depth }
	if myTurn == g.GetPlayerTurn() && g.CheckLose() { return depth -10 }

	if myTurn != g.GetPlayerTurn() && g.CheckWin() { return depth -10 }
	if myTurn != g.GetPlayerTurn() && g.CheckLose() { return 10 - depth }

	return 0
}
