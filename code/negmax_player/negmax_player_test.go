package negmax_player

import(
	"testing"
	"code/game"
)

func TestSelectPosition(t *testing.T){
	var player = CreateNegMaxPlayer()
	var g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 1)
	g.PlaceTip(1, 1)

	g.UpdateToNextTurn()

	row, col := player.SelectPosition(g)

	if 2 != row{ t.Fatal("faied test") }
	if 1 != col{ t.Fatal("faied test") }
}

func TestBestPosition(t *testing.T){
	var player = CreateNegMaxPlayer()
	var g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 1)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 1)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)
	
	g.UpdateToNextTurn()

	player.negmax(g, 0)

	if 2 != player.bestPosition[0]{ t.Fatal("faied test") }
	if 0 != player.bestPosition[1]{ t.Fatal("faied test") }

	g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 1)
	g.PlaceTip(1, 1)

	g.UpdateToNextTurn()

	player.negmax(g, 0)

	if 2 != player.bestPosition[0]{ t.Fatal("faied test") }
	if 1 != player.bestPosition[1]{ t.Fatal("faied test") }
}

func TestnegmaxDepth(t *testing.T){
	var player = CreateNegMaxPlayer()

	// win
	var g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 1)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 1)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)
	
	g.UpdateToNextTurn()

	if 8 != player.negmax(g, 0) { t.Fatal("faied test") }

	// draw
	g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(1, 2)
	g.PlaceTip(2, 0)

	g.UpdateToNextTurn()
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 0)
	g.PlaceTip(1, 1)
	g.PlaceTip(2, 1)

	g.UpdateToNextTurn()

	if 0 != player.negmax(g, 0) { t.Fatal("faied test") }

	// lose
	g = game.CreateNewGame()

	g.PlaceTip(0, 1)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 1)

	if -8 != player.negmax(g, 0) { t.Fatal("faied test") }

	g = game.CreateNewGame()
	if 0 != player.negmax(g, 0) { t.Fatal("faied test") }
}

func TestnegmaxWhenEnd(t *testing.T){
	var player = CreateNegMaxPlayer()

	// win
	var g = game.CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)

	if 10 != player.negmax(g, 0) { t.Fatal("faied test") }

	// draw
	g = game.CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(1, 2)
	g.PlaceTip(2, 0)
	g.PlaceTip(2, 2)

	g.UpdateToNextTurn()
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 0)
	g.PlaceTip(1, 1)
	g.PlaceTip(2, 1)

	if 0 != player.negmax(g, 0) { t.Fatal("faied test") }

	// lose
	g = game.CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)
	g.UpdateToNextTurn()
	if -10 != player.negmax(g, 0) { t.Fatal("faied test") }

}

func TestNegMaxPlayer(t *testing.T){
	var g = game.CreateNewGame()
	var score = evaluate(g, 0)

	if score != 0 { t.Fatal("faied test") }

	score = evaluate(g, 5)
	if score != 0 { t.Fatal("faied test") }

	g = game.CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)

	score = evaluate(g, 6)
	if score != 4 { t.Fatal("faied test") }

	g.UpdateToNextTurn()
	score = evaluate(g, 6)
	if score != -4 { t.Fatal("faied test") }

}

func TestGetMaxIdx(t *testing.T){
	var slice = [] int{1, 2, 3}
	if 2 != getMaxIdx(slice) { t.Fatal("faied test") }

	slice = [] int{1}
	if 0 != getMaxIdx(slice) { t.Fatal("faied test") }
}