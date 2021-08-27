package alphabeta_player

import(
	"testing"
	"code/game"
)

func TestSelectPosition(t *testing.T){
	var player = CreateAlphabetaPlayer()
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
	var player = CreateAlphabetaPlayer()
	var g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 1)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 1)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)
	
	g.UpdateToNextTurn()

	player.alphabeta(g, 0, -9999999, 9999999)

	if 2 != player.bestPosition[0]{ t.Fatal("faied test") }
	if 0 != player.bestPosition[1]{ t.Fatal("faied test") }

	player = CreateAlphabetaPlayer()
	g = game.CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 1)
	g.PlaceTip(1, 1)

	g.UpdateToNextTurn()

	player.alphabeta(g, 0, -9999999, 9999999)

	if 2 != player.bestPosition[0]{ t.Fatal("faied test") }
	if 1 != player.bestPosition[1]{ t.Fatal("faied test") }
}

func TestalphabetaDepth(t *testing.T){
	var player = CreateAlphabetaPlayer()

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

	if 8 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }

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

	if 0 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }

	// lose
	g = game.CreateNewGame()

	g.PlaceTip(0, 1)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)

	g.UpdateToNextTurn()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 1)

	if -8 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }

	g = game.CreateNewGame()
	if 0 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }
}

func TestalphabetaWhenEnd(t *testing.T){
	var player = CreateAlphabetaPlayer()

	// win
	var g = game.CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)

	if 10 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }

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

	if 0 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }

	// lose
	g = game.CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)
	g.UpdateToNextTurn()
	if -10 != player.alphabeta(g, 0, -9999999, 9999999) { t.Fatal("faied test") }

}

func TestAlphabetaPlayer(t *testing.T){
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
