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
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if err := test.Place(JWT_Cookie, sudoku.Placement{Row: i, Column: j, Value: 1}); err != nil {
				t.Fail()
			}

		}
	}
}
