package game

import(
	"testing"
)

func TestCreateNewGame(t *testing.T){
	g := CreateNewGame()

	if g.playerTurn != 0{  t.Fatal("failed test") }

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

	if g.playerTurn != 0{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.playerTurn != 1{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.playerTurn != 0{  t.Fatal("failed test") }
}

func TestgetCurrentTip(t *testing.T){
	g := CreateNewGame()

	if g.getCurrentPlayerTip() != BLACK_TIP{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.getCurrentPlayerTip() != WHITE_TIP{  t.Fatal("failed test") }
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

func TestcanUpdateToNextTurn(t *testing.T){
	var g = CreateNewGame()

	if !g.canUpdateToNextTurn(){  t.Fatal("failed test") }

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)

	if !g.canUpdateToNextTurn(){  t.Fatal("failed test") }

	g.PlaceTip(1, 0)
	g.PlaceTip(1, 1)
	g.PlaceTip(1, 2)

	if !g.canUpdateToNextTurn(){  t.Fatal("failed test") }

	g.PlaceTip(2, 0)
	g.PlaceTip(2, 1)
	g.PlaceTip(2, 2)

	if g.canUpdateToNextTurn(){  t.Fatal("failed test") }
}

func TestCheckDraw(t *testing.T){
	var g = CreateNewGame()

	if g.CheckDraw(){ t.Fatal("failed test") }

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

	if !g.CheckDraw(){ t.Fatal("failed test") }
}

func TestCheckEnd(t *testing.T){
	var g = CreateNewGame()

	if g.CheckEnd(){ t.Fatal("failed test") }

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)

	if !g.CheckEnd(){ t.Fatal("failed test") }

	g = CreateNewGame()

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

	if !g.CheckEnd(){ t.Fatal("failed test") }
}

func TestGetMatchResult(t *testing.T){
	var g = CreateNewGame()

	g.PlaceTip(0, 0)
	g.PlaceTip(0, 1)
	g.PlaceTip(0, 2)

	if WIN != g.GetMatchResult(BLACK){ t.Fatal("failed test") }
	if LOSE != g.GetMatchResult(WHITE){ t.Fatal("failed test") }
}

func TestGetMatchResultWhenDraw(t *testing.T){
	var g = CreateNewGame()

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

	if DRAW != g.GetMatchResult(WHITE){ t.Fatal("failed test") }
	if DRAW != g.GetMatchResult(BLACK){ t.Fatal("failed test") }
}