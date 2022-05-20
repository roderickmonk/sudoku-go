package sudoku

type Board [BOARD_SIZE][BOARD_SIZE]byte

type Placement struct {
	Row    int
	Column int
	Value  byte
}


