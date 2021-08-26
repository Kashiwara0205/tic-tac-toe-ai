package random_player

import(
	"time"
	"code/game"
	"math/rand"
)

type Randomizer struct{}
func (r *Randomizer) createValue() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

type RandomizerInterface interface{
	createValue() int
}

type RandomPlayer struct{ 
	randomizer RandomizerInterface
}

func CreateRandomPlayer() RandomPlayer{
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
