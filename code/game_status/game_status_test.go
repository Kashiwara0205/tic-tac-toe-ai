package game_status

import(
	"code/game"
	"testing"
)

func TestCreateGameStatus(t *testing.T){
	g := game.CreateNewGame()
	status := CreateGameStatus(g)

	if status.playerTurn != 0{  t.Fatal("failed test") }

	if status.board[0][0] != 0{ t.Fatal("failed test") }
	if status.board[0][1] != 0{ t.Fatal("failed test") }
	if status.board[0][2] != 0{ t.Fatal("failed test") }

	if status.board[1][0] != 0{ t.Fatal("failed test") }
	if status.board[1][1] != 0{ t.Fatal("failed test") }
	if status.board[1][2] != 0{ t.Fatal("failed test") }

	if status.board[2][0] != 0{ t.Fatal("failed test") }
	if status.board[2][1] != 0{ t.Fatal("failed test") }
	if status.board[2][2] != 0{ t.Fatal("failed test") }
}
