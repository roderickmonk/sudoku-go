package sudoku

type Board [9][9]byte

type Placement struct {
	I     byte
	J     byte
	Value byte
}
