package main

import (
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
)

func Test_EmptyBoard(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	// Set an empty board
	test.SetBoard(JWT_Cookie, &sudoku.Board{})

	// Ensure all board positions are available
	for row := 0; row < sudoku.BOARD_SIZE; row++ {
		for column := 0; column < sudoku.BOARD_SIZE; column++ {
			if err := test.Place(JWT_Cookie, sudoku.Placement{Row: row, Column: column, Value: 1}); err == nil {
				t.Fail() // A successful placement is a test failure
			}
		}
	}
}
