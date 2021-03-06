package sudoku

// Derived from git@github.com:dsparling/go-sudoku-grid-generator.git

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeGame() Board {
	rand.Seed(time.Now().UTC().UnixNano())

	board := Board{}
	board[0] = [BOARD_SIZE]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	board[1] = [BOARD_SIZE]byte{4, 5, 6, 7, 8, 9, 1, 2, 3}
	board[2] = [BOARD_SIZE]byte{7, 8, 9, 1, 2, 3, 4, 5, 6}
	board[3] = [BOARD_SIZE]byte{2, 3, 4, 5, 6, 7, 8, 9, 1}
	board[4] = [BOARD_SIZE]byte{5, 6, 7, 8, 9, 1, 2, 3, 4}
	board[5] = [BOARD_SIZE]byte{8, 9, 1, 2, 3, 4, 5, 6, 7}
	board[6] = [BOARD_SIZE]byte{3, 4, 5, 6, 7, 8, 9, 1, 2}
	board[7] = [BOARD_SIZE]byte{6, 7, 8, 9, 1, 2, 3, 4, 5}
	board[8] = [BOARD_SIZE]byte{9, 1, 2, 3, 4, 5, 6, 7, 8}

	createSolvedSudoku(&board)
	PrintBoard(&board)
	createUnsolvedSudoku(&board)
	PrintBoard(&board)
	// solve(&board)
	// PrintBoard(&board)

	return board
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

/* Randomly swap rows and columns within the 3x3 borders
 * and swap individual numbers globally. */
func createSolvedSudoku(b *Board) {
	for i := 0; i < 100; i++ {
		myrand := randInt(0, 100)
		if myrand < 33 {
			swapLine(b)
		} else if myrand < 66 {
			swapCol(b)
		} else if myrand < 100 {
			swapNumber(b)
		}
	}
}

/* Erase between 50% and 80% of the numbers to create
 * the unsolved grid with a random difficulty. */
func createUnsolvedSudoku(b *Board) {
	r1 := randInt(0, 100) // 0-99
	r2 := randInt(50, 80) // 50-79
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if r1 < r2 {
				b[i][j] = 0
			}
			r1 = randInt(0, 100)
		}
	}
}

func swapLine(b *Board) {
	randInts := [3]int{0, 3, 6}
	r := randInts[rand.Intn(len(randInts))]
	l1 := r + randInt(0, 3) // 0-2
	l2 := r + randInt(0, 3) // 0-2
	b[l1], b[l2] = b[l2], b[l1]
}

func swapCol(b *Board) {
	randInts := [3]int{0, 3, 6}
	r := randInts[rand.Intn(len(randInts))]
	c1 := r + randInt(0, 3) // 0-2
	c2 := r + randInt(0, 3) // 0-2
	for line := 0; line < 9; line++ {
		b[line][c1], b[line][c2] = b[line][c2], b[line][c1]
	}
}

func swapNumber(b *Board) {
	n1 := byte(randInt(1, 10))
	n2 := byte(randInt(1, 10))

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if b[i][j] == n1 {
				b[i][j] = n2
			} else if b[i][j] == n2 {
				b[i][j] = n1
			}
		}
	}
}

/* Simple solving algorithm:
 * 1.List all empty cells in order.
 * 2.Take the first empty cell as current cell.
 * 3.Fill the current cell with the current cell value +1.
 * 4.If this number violate the Sudoku condition back to (3), if not
 * back to (3) with the next cell as current cell.If the number == 9 */
func solve(b *Board) {
	var list []int
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if b[j][i] == 0 {
				list = append(list, (j*10)+i)
			}
		}
	}

	it := 0
	for i := 0; i < len(list); i++ {
		t := byte(0)
		l := list[i]
		jc := l % 10
		ic := (l - l%10) / 10

		for {
			t = b[ic][jc] + 1
			b[ic][jc] = t
			it++
			if verify(l, b) {
				break
			}
		}

		if t > BOARD_SIZE {
			if i == len(list) {
				i = len(list) + 1
			} else {
				b[ic][jc] = 0
				if i > 0 {
					i = i - 2
				} else {
					i = -1
				}
			}
		} else if l == 88 {
			i = len(list) + 1
		}
	}

	fmt.Printf("Number of iterations: %d\n", it)
	fmt.Print("Difficulty rating: ")
	if it < 1000 {
		fmt.Println("Easy")
	} else if it < 10000 {
		fmt.Println("Normal")
	} else if it < 100000 {
		fmt.Println("Hard")
	} else if it < 1000000 {
		fmt.Println("Very hard")
	} else if it < 10000000 {
		fmt.Println("Extreme")
	}
	fmt.Println("\nSolution")
}

func verify(l int, b *Board) bool {
	j := l % 10
	i := (l - l%10) / 10
	var ic, jc int

	boolVal := true
	for i1 := 0; i1 < 9; i1++ {
		if i1 != i && b[i1][j] == b[i][j] {
			boolVal = false
		}
	}
	for j1 := 0; j1 < 9; j1++ {
		if j1 != j && b[i][j1] == b[i][j] {
			boolVal = false
		}
	}

	if i >= 0 && i < 3 {
		ic = 1
	} else if i > 2 && i < 6 {
		ic = 2
	} else {
		ic = 3
	}

	if j >= 0 && j < 3 {
		jc = 1
	} else if j > 2 && j < 6 {
		jc = 2
	} else {
		jc = 3
	}

	for i1 := ic*3 - 3; i1 < ic*3; i1++ {
		for j1 := jc*3 - 3; j1 < jc*3; j1++ {
			if (j1 != j || i1 != i) && b[i1][j1] == b[i][j] {
				boolVal = false
			}
		}
	}
	return boolVal
}
