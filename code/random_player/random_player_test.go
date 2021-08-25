package random_player

import(
	"code/game"
	"code/game_status"
	"testing"
)

type MockRandomizer struct{}
func (r *MockRandomizer) createValue() int{
	return 1
}

func TestSelectPosition(t *testing.T){
	var player = createRandomPlayer()

	g := game.CreateNewGame()
	var status = game_status.CreateGameStatus(g)

	var row, col = player.SelectPosition(status)
	
	if row != 0{ t.Fatal("faied test") }
	if col != 0{ t.Fatal("faied test") }

	randomizer := &MockRandomizer{}
	player = RandomPlayer{ randomizer: randomizer }

	g.PlaceTip(0, 0)

	status = game_status.CreateGameStatus(g)
	row, col = player.SelectPosition(status)
	
	if row != 1{ t.Fatal("faied test") }
	if col != 1{ t.Fatal("faied test") }
}
