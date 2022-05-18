package sudoku

import "fmt"

func PrintBoard(b *Board) {
	
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(b[i][j], " ")
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}