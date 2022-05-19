package main

import (
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
)

func Test_ColumnConflict(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	for i := byte(0); i < 9; i++ {
		for j := byte(0); j < 9; j++ {

			board := sudoku.Board{}
			board[i][j] = 1
			test.SetBoard(JWT_Cookie, &board)

			if err := test.Place(JWT_Cookie, sudoku.Placement{I: 8, J: j, Value: 1}); err != nil {
				continue
			} else {
				t.Fail()
			}
		}
	}
}
