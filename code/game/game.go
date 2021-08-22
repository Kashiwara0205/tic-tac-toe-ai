package game

const (
	NONE = iota
	BLACK_TIP
	WHITE_TIP
)

const (
	BLACK_TURN = iota
	WHITE_TURN
)

type Game struct{
	board [3][3]int
	turn int
}

func CreateNewGame() Game{
	board := [3][3]int{
		{NONE, NONE, NONE},
		{NONE, NONE, NONE},
		{NONE, NONE, NONE},
	}

	return Game{ turn: BLACK_TURN, board: board }
}

func (g *Game) UpdateToNextTurn(){
	if g.turn == 0{
		g.turn = 1
		return 
	}

	if g.turn == 1{
		g.turn = 0
	}
}

func (g *Game) getTipStatus(row int, col int) int{
	return g.board[row][col]
}

func (g *Game) PlaceTip(row int, col int){
	g.board[row][col] = g.getCurrentTip()
}

func (g *Game) CanPlaceTip(row int, col int) bool{
	if (row > 2 || row < 0){ return false }
	if (col > 2 || col < 0){ return false }

	tip := g.getTipStatus(row, col)

	if (tip == NONE){ return true }
	return false
}

func (g *Game) CheckWin() bool{
	tip := g.getCurrentTip()

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

func (g *Game) getCurrentTip() int {
	if (BLACK_TURN == g.turn){ 
		return BLACK_TIP 
	}else{ 
		return WHITE_TIP 
	}
}