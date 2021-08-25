package random_player

import(
	"code/game"
	"testing"
)

type MockRandomizer struct{}
func (r *MockRandomizer) createValue() int{
	return 1
}

func TestSelectPosition(t *testing.T){
	var player = createRandomPlayer()

	g := game.CreateNewGame()
	var row, col = player.SelectPosition(g)
	
	if row != 0{ t.Fatal("faied test") }
	if col != 0{ t.Fatal("faied test") }

	randomizer := &MockRandomizer{}
	player = RandomPlayer{ randomizer: randomizer }

	g.PlaceTip(0, 0)
	row, col = player.SelectPosition(g)
	
	if row != 1{ t.Fatal("faied test") }
	if col != 1{ t.Fatal("faied test") }
}
