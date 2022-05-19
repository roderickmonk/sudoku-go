package main

import (
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
)

func Test_CellConflict(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	for i := byte(0); i < 9; i++ {
		for j := byte(0); j < 9; j++ {

            // Ensure a conflict cell placement for all cell positions
            board := sudoku.Board{}
            board[i][j]=1
            test.SetBoard(JWT_Cookie, &board)

			if err := test.Place(JWT_Cookie, sudoku.Placement{I: i, J: j, Value: 1}); err != nil {
				continue
			} else {
                t.Fail()
            }
		}
	}
}
