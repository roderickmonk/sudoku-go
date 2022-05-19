package main

import (
	"fmt"
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
)

func Test_EmptyBoard(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	// Set an empty board
	test.SetBoard(JWT_Cookie, &sudoku.Board{})

	// Ensure all board positions are available
	for i := byte(0); i < 9; i++ {
		for j := byte(0); j < 9; j++ {
			if err := test.Place(JWT_Cookie, sudoku.Placement{I: i, J: j, Value: 1}); err != nil {
				fmt.Printf("Failure: %v\n", err)
				t.Fail()
			}

		}
	}
}
