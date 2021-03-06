package game

import (
	"fmt"
)

const (
	NONE = iota
	BLACK_TIP
	WHITE_TIP
)

const (
	BLACK = iota
	WHITE
)

const (
	WIN = iota
	LOSE
	DRAW
)

type Game struct{
	board [3][3]int
	playerTurn int
}

func CreateNewGame() Game{
	board := [3][3]int{
		{NONE, NONE, NONE},
		{NONE, NONE, NONE},
		{NONE, NONE, NONE},
	}

	return Game{ playerTurn: BLACK, board: board }
}

func (g *Game) CopyGame() Game{
	return Game{ playerTurn: g.GetPlayerTurn(), board: g.GetBoard() }
}

func (g *Game) GetBoard() [3][3]int{
	return g.board
}

func (g *Game) GetPlayerTurn() int{
	return g.playerTurn
}

func (g *Game) UpdateToNextTurn(){
	if g.playerTurn == 0{
		g.playerTurn = 1
		return 
	}

	if g.playerTurn == 1{
		g.playerTurn = 0
	}
}

func (g *Game) getTipStatus(row int, col int) int{
	return g.board[row][col]
}

func (g *Game) PlaceTip(row int, col int){
	g.board[row][col] = g.getCurrentPlayerTip()
}

func (g *Game) CanPlaceTip(row int, col int) bool{
	if (row > 2 || row < 0){ return false }
	if (col > 2 || col < 0){ return false }

	tip := g.getTipStatus(row, col)

	if (tip == NONE){ return true }
	return false
}

func (g *Game) CheckEnd() bool{
	if g.CheckWin() { return true }
	if g.CheckLose() { return true }
	if !g.canUpdateToNextTurn() { return true }

	return false
}

func (g *Game) canUpdateToNextTurn() bool{
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if ( NONE == g.getTipStatus(i, j) ){
				return true
			}
		}
	}

	return false
}

func (g *Game) CheckDraw() bool{
	return ( !g.canUpdateToNextTurn() && !g.checkWin(BLACK_TIP) && !g.checkWin(WHITE_TIP) )
}

func (g *Game) CheckWin() bool {
	return g.checkWin(g.getCurrentPlayerTip())
}

func (g *Game) CheckLose() bool {
	return g.checkWin(g.getOpponentPlayerTip())
}

func (g *Game) checkWin(tip int) bool{
	if ( ( tip == g.getTipStatus(0, 0) ) && ( tip == g.getTipStatus(0, 1) ) && ( tip == g.getTipStatus(0, 2) ) ){ return true }
	if ( ( tip == g.getTipStatus(1, 0) ) && ( tip == g.getTipStatus(1, 1) ) && ( tip == g.getTipStatus(1, 2) ) ){ return true }
	if ( ( tip == g.getTipStatus(2, 0) ) && ( tip == g.getTipStatus(2, 1) ) && ( tip == g.getTipStatus(2, 2) ) ){ return true }

	if ( ( tip == g.getTipStatus(0, 0) ) && ( tip == g.getTipStatus(1, 0) ) && ( tip == g.getTipStatus(2, 0) ) ){ return true }
	if ( ( tip == g.getTipStatus(0, 1) ) && ( tip == g.getTipStatus(1, 1) ) && ( tip == g.getTipStatus(2, 1) ) ){ return true }
	if ( ( tip == g.getTipStatus(0, 2) ) && ( tip == g.getTipStatus(1, 2) ) && ( tip == g.getTipStatus(2, 2) ) ){ return true }

	if ( ( tip == g.getTipStatus(0, 0) ) && ( tip == g.getTipStatus(1, 1) ) && ( tip == g.getTipStatus(2, 2) ) ){ return true }
	if ( ( tip == g.getTipStatus(0, 2) ) && ( tip == g.getTipStatus(1, 1) ) && ( tip == g.getTipStatus(2, 0) ) ){ return true }

	return false
}

func  (g *Game) getCurrentPlayerTip() int {
	return g.getTip(g.playerTurn)
}

func  (g *Game) getOpponentPlayerTip() int {
	var opponent = 0
	if (BLACK == g.playerTurn){ 
		opponent = WHITE
	}else{ 
		opponent = BLACK 
	}

	return g.getTip(opponent)
}

func (g *Game) getTip(player int) int {
	if (BLACK == player){ 
		return BLACK_TIP 
	}else{ 
		return WHITE_TIP 
	}
}

func (g *Game) GetMatchResult(player int) int {
	if (g.CheckDraw()){ return DRAW }

	tip := g.getTip(player)
	if (g.checkWin(tip)){
		return WIN
	}else{
		return LOSE
	}
}

func (g *Game) PrintBoard(){
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{

			if ( NONE == g.getTipStatus(i, j) ){
				fmt.Printf(" N ")
			}

			if ( BLACK_TIP == g.getTipStatus(i, j) ){
				fmt.Printf(" x ")
			}

			if ( WHITE_TIP == g.getTipStatus(i, j) ){
				fmt.Printf(" o ")
			}
		}
		fmt.Printf("\n")
	}
}
