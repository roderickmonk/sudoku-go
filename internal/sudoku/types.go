package sudoku

type Board [9][9]byte

type Placement struct {
	Row    int
	Column int
	Value  byte
}


