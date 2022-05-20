package main

import (

	// "fmt"
	"net/http"
	"testing"

	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
	"github.com/stretchr/testify/assert"
)

func Test_General(t *testing.T) {

	var (
		err        error
		JWT_Cookie *http.Cookie
	)

	if JWT_Cookie, _, err = test.SignIn(t); err != nil {
		assert.FailNow(t, err.Error())
		t.Fail()
	}


	if err = test.Place(JWT_Cookie, sudoku.Placement{Row: 0, Column: 1, Value: 4}); err != nil {
		assert.FailNow(t, err.Error())
		t.Fail()
	}

	if _, err = test.Refresh(JWT_Cookie); err != nil {
		assert.FailNow(t, err.Error())
		t.Fail()
	}

	// Set an empty board
	test.SetBoard(JWT_Cookie, &sudoku.Board{})

}
