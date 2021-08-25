package minmax_player

import(
	"testing"
	"code/game"
)

func TestMinMaxPlayer(t *testing.T){
	g := game.CreateNewGame()
	player := CreateMinMaxPlayer()
	var score = player.evaluate(g)

	if score != 0 { t.Fatal("faied test") }
}
