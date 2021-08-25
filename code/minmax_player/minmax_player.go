package minmax_player

import(
	"code/game"
)

type MinMaxPlayer struct{ }

func CreateMinMaxPlayer() MinMaxPlayer{
	return MinMaxPlayer{}
}

func minxmax(g game.Game, depth int, myTurn int) int {
	if g.CheckEnd(){ return evaluate(g, depth, myTurn) }
	depth += 1

	var scores[]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if ( g.CanPlaceTip(i, j) ){
				possibleGame := g.CopyGame()
				possibleGame.PlaceTip(i, j)
				possibleGame.UpdateToNextTurn()
				scores = append(scores, minxmax(possibleGame, depth, myTurn) )
			}
		}
	}

	if ( myTurn == g.GetPlayerTurn() ){
		maxScoreIdx := getMaxIdx(scores)
		return scores[maxScoreIdx]
	}else{
		minScoreIdx := getMinIdx(scores)
		return scores[minScoreIdx]
	}
}

func getMaxIdx(slice []int) int {
	max := -9999999
	maxIdx := -1
	for idx, s := range slice {
		if max < s {
			max = s
			maxIdx = idx
		}
	}

	return maxIdx
}

func getMinIdx(slice []int) int {
	min := 9999999
	minIdx := -1
	for idx, s := range slice {
		if min > s {
			min = s
			minIdx = idx
		}
	}

	return minIdx
}

func evaluate(g game.Game, depth int, myTurn int) int {
	if myTurn == g.GetPlayerTurn() && g.CheckWin() { return 10 - depth }
	if myTurn == g.GetPlayerTurn() && g.CheckLose() { return depth -10 }

	if myTurn != g.GetPlayerTurn() && g.CheckWin() { return depth -10 }
	if myTurn != g.GetPlayerTurn() && g.CheckLose() { return 10 - depth }

	return 0
}
