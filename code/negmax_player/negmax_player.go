package negmax_player

import(
	"code/game"
	"code/utils"
)

type NegMaxPlayer struct{
	bestPosition [2]int
}

func CreateNegMaxPlayer() NegMaxPlayer{
	return NegMaxPlayer{}
}

func (player *NegMaxPlayer) SelectPosition(g game.Game) (int, int) {
	player.negmax(g, 0)

	row := player.bestPosition[0]
	col := player.bestPosition[1]

	return row, col
}

func (player *NegMaxPlayer) negmax(g game.Game, depth int) int {
	if g.CheckEnd(){ return evaluate(g, depth) }
	depth += 1

	var scores[]int
	var positions[][]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if ( g.CanPlaceTip(i, j) ){
				possibleGame := g.CopyGame()
				possibleGame.PlaceTip(i, j)
				possibleGame.UpdateToNextTurn()
				scores = append(scores, -player.negmax(possibleGame, depth) )
				position := []int{i, j}
				positions = append(positions, position)
			}
		}
	}

	maxScoreIdx := utils.GetMaxIdx(scores)
	player.bestPosition[0] = positions[maxScoreIdx][0]
	player.bestPosition[1] = positions[maxScoreIdx][1]

	return scores[maxScoreIdx]
}

func evaluate(g game.Game, depth int) int {
	if g.CheckWin() { return 10 - depth }
	if g.CheckLose() { return depth -10 }

	return 0
}
