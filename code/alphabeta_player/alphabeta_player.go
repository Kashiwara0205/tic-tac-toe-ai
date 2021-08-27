package alphabeta_player

import(
	"code/game"
)

type AlphabetaPlayer struct{
	bestScore int
	bestPosition [2]int
}

func CreateAlphabetaPlayer() AlphabetaPlayer{
	return AlphabetaPlayer{}
}

func (player *AlphabetaPlayer) SelectPosition(g game.Game) (int, int) {
	player.alphabeta(g, 0, 0, 9999999)

	row := player.bestPosition[0]
	col := player.bestPosition[1]

	return row, col
}

func (player *AlphabetaPlayer) alphabeta(g game.Game, depth int, alpha int , beta int) int {
	if g.CheckEnd(){ return evaluate(g, depth) }
	depth += 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if ( g.CanPlaceTip(i, j) ){
				possibleGame := g.CopyGame()
				possibleGame.PlaceTip(i, j)
				possibleGame.UpdateToNextTurn()
				score := -player.alphabeta(possibleGame, depth, -beta, -alpha)
				if score > alpha{
					alpha = score
					player.bestScore = score
					player.bestPosition[0] = i
					player.bestPosition[1] = j
					
					if alpha >= beta{
						return player.bestScore
					}
				}
			}
		}
	}

	return player.bestScore
}

func evaluate(g game.Game, depth int) int {
	if g.CheckWin() { return 10 - depth }
	if g.CheckLose() { return depth -10 }

	return 0
}
