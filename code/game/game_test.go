package game

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewGame(t *testing.T){
	g := CreateNewGame()

	if g.turn != 0{  t.Fatal("failed test") }

	if g.GetTipStatus(0, 0) != 0{ t.Fatal("failed test") }
	if g.GetTipStatus(0, 1) != 0{ t.Fatal("failed test") }
	if g.GetTipStatus(0, 2) != 0{ t.Fatal("failed test") }

	if g.GetTipStatus(1, 0) != 0{ t.Fatal("failed test") }
	if g.GetTipStatus(1, 1) != 0{ t.Fatal("failed test") }
	if g.GetTipStatus(1, 2) != 0{ t.Fatal("failed test") }

	if g.GetTipStatus(2, 0) != 0{ t.Fatal("failed test") }
	if g.GetTipStatus(2, 1) != 0{ t.Fatal("failed test") }
	if g.GetTipStatus(2, 2) != 0{ t.Fatal("failed test") }
}

func TestGetTipStatusRange(t *testing.T){
	g := CreateNewGame()

	if g.GetTipStatus(0, 0) != 0{ t.Fatal("faied test") }
	assert.Panics(t, func() { g.GetTipStatus(-1, 0) }, "The code did not panic")

	if g.GetTipStatus(2, 0) != 0{ t.Fatal("failed test") }
	assert.Panics(t, func() { g.GetTipStatus(3, 0) }, "The code did not panic")

	if g.GetTipStatus(0, 0) != 0{ t.Fatal("faied test") }
	assert.Panics(t, func() { g.GetTipStatus(0, -1) }, "The code did not panic")

	if g.GetTipStatus(0, 2) != 0{ t.Fatal("failed test") }
	assert.Panics(t, func() { g.GetTipStatus(0, 3) }, "The code did not panic")
}

func TestUpdateToNextTurn(t *testing.T){
	g := CreateNewGame()

	if g.turn != 0{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.turn != 1{  t.Fatal("failed test") }

	g.UpdateToNextTurn()

	if g.turn != 0{  t.Fatal("failed test") }
}