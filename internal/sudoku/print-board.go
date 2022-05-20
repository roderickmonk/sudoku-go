package sudoku

import "fmt"

func PrintBoard(b *Board) {
	
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			fmt.Print(b[i][j], " ")
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}