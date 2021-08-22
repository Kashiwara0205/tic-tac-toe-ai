package game

const (
	NONE = iota
	BLACK_TIP
	WHITE_TIP
)

type Game struct{
	board [3][3]int
}

func CreateNewGame() Game{
	board := [3][3]int{
		{NONE, NONE, NONE},
		{NONE, NONE, NONE},
		{NONE, NONE, NONE},
	}

	return Game{ board: board }
}

func (g *Game) GetTipStatus(row int, col int) int{
	if (row > 2 || row < 0){ panic("row is out of range") }
	if (col > 2 || col < 0){ panic("col is out of range") }

	return g.board[row][col]
}