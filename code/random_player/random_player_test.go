package random_player

import(
	"testing"
)

type MockRandomizer struct{}
func (r *MockRandomizer) createValue() int{
	return 1
}

func TestSelectPosition(t *testing.T){
	var player = createRandomPlayer()

	var board = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	var row, col = player.SelectPosition(board)
	
	if row != 0{ t.Fatal("faied test") }
	if col != 0{ t.Fatal("faied test") }


	randomizer := &MockRandomizer{}
	player = RandomPlayer{ randomizer: randomizer }

	board = [3][3]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	row, col = player.SelectPosition(board)
	
	if row != 1{ t.Fatal("faied test") }
	if col != 1{ t.Fatal("faied test") }
}
