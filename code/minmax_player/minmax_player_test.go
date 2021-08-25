package minmax_player

import(
	"testing"
	"code/game"
)

func TestMinMaxPlayer(t *testing.T){
	g := game.CreateNewGame()
	var score = evaluate(g, 0, 0)

	if score != 0 { t.Fatal("faied test") }
}

func TestgetMaxIdx(t *testing.T){
	var slice = [] int{1, 2, 3}
	if 2 != getMaxIdx(slice) { t.Fatal("faied test") }

	slice = [] int{1}
	if 0 != getMaxIdx(slice) { t.Fatal("faied test") }
}

func TestgetMinIdx(t *testing.T){
	var slice = [] int{1, 2, 3}
	if 0 != getMaxIdx(slice) { t.Fatal("faied test") }

	slice = [] int{3, 2, 1}
	if 2 != getMaxIdx(slice) { t.Fatal("faied test") }
}