package random_player

import(
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

func (player *RandomPlayer) SelectPosition(board [3][3]int) (int, int) {
	var row = 0
	var col = 0

	for 0 != board[row][col]{
		row = player.randomizer.createValue()
		col = player.randomizer.createValue()
	}

	return row, col
}
