package random_player

import(
	"code/game"
	"math/rand"
)

type Randomizer struct{}
func (r *Randomizer) createValue() int{
	return rand.Intn(3)
}

type RandomizerInterface interface{
	createValue() int
}

type RandomPlayer struct{ 
	randomizer RandomizerInterface
}

func createRandomPlayer() RandomPlayer{
	randomizer := &Randomizer{}
	return RandomPlayer{ randomizer: randomizer }
}

func (player *RandomPlayer) SelectPosition(g game.Game) (int, int) {
	var row = 0
	var col = 0

	for !g.CanPlaceTip(row, col){
		row = player.randomizer.createValue()
		col = player.randomizer.createValue()
	}

	return row, col
}
