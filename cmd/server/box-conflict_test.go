package main

import (
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
)

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func Test_BoxConflict(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	boxes := [][][]int{
		// Top boxes
		{{0, 1, 2}, {0, 1, 2}},
		{{0, 1, 2}, {3, 4, 5}},
		{{0, 1, 2}, {6, 7, 8}},
		// middle boxes
		{{3, 4, 5}, {0, 1, 2}},
		{{3, 4, 5}, {3, 4, 5}},
		{{3, 4, 5}, {6, 7, 8}},
		// bottom boxes
		{{6, 7, 8}, {0, 1, 2}},
		{{6, 7, 8}, {3, 4, 5}},
		{{6, 7, 8}, {6, 7, 8}}}

	for _, box := range boxes {

		rows := box[0]
		cols := box[1]

		for _, row := range rows {

			for _, col := range cols {

				board := sudoku.Board{}
				board[row][col] = 1
				test.SetBoard(JWT_Cookie, &board)

				// Find some other cell in the same box
				conflict_row := If(row == rows[0], rows[1], If(row == rows[1], rows[2], rows[0]))
				conflict_col := If(col == cols[0], cols[1], If(col == cols[1], cols[2], cols[0]))

				// The placement must always fail
				if err := test.Place(JWT_Cookie, sudoku.Placement{Row: conflict_row, Column: conflict_col, Value: 1}); err != nil {
					continue
				} else {
					t.Fail() // A successful placement is a test failure
				}
			}
		}

	}
}
