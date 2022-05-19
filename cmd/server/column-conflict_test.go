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

/*
for (let i = 0; i < 8; ++i) {
	for (let j = 0; j < 9; ++j) {

		const board: Board = Array(81).fill(null);
		board[9 * i + j] = 4;
		await sudoku.setBoard(board);

		try {
			await sudoku.place({ i: 8, j, value: 4 });
			throw new Error(`Illegal placement (${i},${j})`);
		} catch (err) {
			// expecting 403s
			if (err.response.status !== 403 ||
				![PlaceResult.ColumnConflict, PlaceResult.BoxConflict].includes(err.response.data)) {
				console.log({ responseData: err.response.data });
				throw err;
			}
		}
	}
}
console.log("Column Conflict Testing Successful");
*/
