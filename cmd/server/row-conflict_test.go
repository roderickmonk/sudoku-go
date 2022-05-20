package main

import (
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
)

func Test_RowConflict(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	for i := 0; i < sudoku.BOARD_SIZE; i++ {
		for j := 0; j < sudoku.BOARD_SIZE; j++ {

			board := sudoku.Board{}
			board[i][j] = 1
			test.SetBoard(JWT_Cookie, &board)

			if err := test.Place(JWT_Cookie, sudoku.Placement{Row: i, Column: 8, Value: 1}); err == nil {
				t.Fail() // A successful placement is a test failure
			}
		}
	}
}
