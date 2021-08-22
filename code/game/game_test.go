package game

import(
	"testing"
)

func TestCreateNewGame(t *testing.T){
	g := CreateNewGame()

	if g.turn != 0{  t.Fatal("failed test") }

	if g.getTipStatus(0, 0) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(0, 1) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(0, 2) != 0{ t.Fatal("failed test") }

	if g.getTipStatus(1, 0) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(1, 1) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(1, 2) != 0{ t.Fatal("failed test") }

	if g.getTipStatus(2, 0) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(2, 1) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(2, 2) != 0{ t.Fatal("failed test") }
}

func TestgetTipStatusRange(t *testing.T){
	g := CreateNewGame()

	if g.getTipStatus(0, 0) != 0{ t.Fatal("faied test") }
	if g.getTipStatus(2, 0) != 0{ t.Fatal("failed test") }
	if g.getTipStatus(0, 0) != 0{ t.Fatal("faied test") }
	if g.getTipStatus(0, 2) != 0{ t.Fatal("failed test") }
}

func TestUpdateToNextTurn(t *testing.T){
	g := CreateNewGame()

	if g.turn != 0{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.turn != 1{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.turn != 0{  t.Fatal("failed test") }
}

func TestgetCurrentTip(t *testing.T){
	g := CreateNewGame()

	if g.getCurrentTip() != BLACK_TIP{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.getCurrentTip() != WHITE_TIP{  t.Fatal("failed test") }
}

func TestPlaceTip(t *testing.T){
	g := CreateNewGame()

	g.PlaceTip(0, 0)

	if g.getTipStatus(0, 0) != BLACK_TIP{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	g.PlaceTip(2, 2)

	if g.getTipStatus(2, 2) != WHITE_TIP{  t.Fatal("failed test") }
}

func TestCanPlaceTip(t *testing.T){
	g := CreateNewGame()

	if !g.CanPlaceTip(0, 0){  t.Fatal("failed test") }
	if !g.CanPlaceTip(0, 1){  t.Fatal("failed test") }
	if !g.CanPlaceTip(0, 2){  t.Fatal("failed test") }

	if !g.CanPlaceTip(1, 0){  t.Fatal("failed test") }
	if !g.CanPlaceTip(1, 1){  t.Fatal("failed test") }
	if !g.CanPlaceTip(1, 2){  t.Fatal("failed test") }

	if !g.CanPlaceTip(2, 0){  t.Fatal("failed test") }
	if !g.CanPlaceTip(2, 1){  t.Fatal("failed test") }
	if !g.CanPlaceTip(2, 2){  t.Fatal("failed test") }

	if g.CanPlaceTip(-1, 0){  t.Fatal("failed test") }
	if g.CanPlaceTip(3, 0){  t.Fatal("failed test") }

	if g.CanPlaceTip(0, -1){  t.Fatal("failed test") }
	if g.CanPlaceTip(0, 3){  t.Fatal("failed test") }

	g.PlaceTip(0, 0)
	g.UpdateToNextTurn()
	g.PlaceTip(0, 1)

	if g.CanPlaceTip(0, 0){  t.Fatal("failed test") }
	if g.CanPlaceTip(0, 1){  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.CanPlaceTip(0, 0){  t.Fatal("failed test") }
	if g.CanPlaceTip(0, 1){  t.Fatal("failed test") }
}

func TestCheckWin(t *testing.T){
	var g = CreateNewGame()
	if g.CheckWin(){  t.Fatal("failed test") }

	g.PlaceTip(0, 0)
	if g.CheckWin(){  t.Fatal("failed test") }
	g.PlaceTip(0, 1)
	if g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(1, 0)
	g.PlaceTip(1, 1)
	g.PlaceTip(1, 2)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(2, 0)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(1, 0)
	g.PlaceTip(2, 0)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(0, 1)
	g.PlaceTip(1, 1)
	g.PlaceTip(2, 1)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 2)
	g.PlaceTip(2, 2)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(0, 2)
	g.PlaceTip(1, 1)
	g.PlaceTip(2, 0)
	if !g.CheckWin(){  t.Fatal("failed test") }

	g = CreateNewGame()
	g.PlaceTip(0, 0)
	g.PlaceTip(1, 1)
	g.PlaceTip(2, 2)
	if !g.CheckWin(){  t.Fatal("failed test") }
}