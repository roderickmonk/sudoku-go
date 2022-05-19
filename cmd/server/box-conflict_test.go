package main

import (
	"github.com/roderickmonk/sudoku-go/internal/sudoku"
	"github.com/roderickmonk/sudoku-go/internal/test"
	"testing"
)

func Test_BoxConflict(t *testing.T) {

	JWT_Cookie, _, _ := test.SignIn(t)

	boxes := [][][]byte{
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
  // Box Conflict Testing
  const boxes = [
	// top boxes
	[[0, 1, 2], [0, 1, 2]],
	[[0, 1, 2], [3, 4, 5]],
	[[0, 1, 2], [6, 7, 8]],
	// middle boxes
	[[3, 4, 5], [0, 1, 2]],
	[[3, 4, 5], [3, 4, 5]],
	[[3, 4, 5], [6, 7, 8]],
	// bottom boxes
	[[6, 7, 8], [0, 1, 2]],
	[[6, 7, 8], [3, 4, 5]],
	[[6, 7, 8], [6, 7, 8]],
];

const value = 4;

for (const [rows, cols] of boxes) {

	for (let i of rows) {
		for (let j of cols) {

			const board: Board = Array(81).fill(null);
			board[9 * i + j] = value;
			await sudoku.setBoard(board);

			// Find some other cell in the same box
			const ii = i === rows[0] ? rows[1] : i === rows[1] ? rows[2] : rows[0];
			const jj = j === cols[0] ? cols[1] : j === cols[1] ? cols[2] : cols[0];

			try {
				await sudoku.place({ i: ii, j: jj, value });
				throw new Error(`Illegal placement (${ii},${jj})`);
			} catch (err) {
				// expecting 403s
				if (err.response.status !== 403 ||
					![PlaceResult.BoxConflict].includes(err.response.data)) {
					console.log({ responseData: err.response.data });
					throw err;
				}
			}
		}
	}
}
console.log("Box Conflict Testing Successful");
*/
