// golang entry for each AlgoExpert challenge
// #[Recursion]/Solve Sudoku

package main

import (
	"fmt"

	recursion "github.com/letempest/algoexpert-in-action/Recursion/Solve-Sudoku"
)

func main() {
	board := [9][9]int{
		{7, 8, 0, 4, 0, 0, 1, 2, 0}, // [7, 8, 5, 4, 3, 9, 1, 2, 6]
		{6, 0, 0, 0, 7, 5, 0, 0, 9}, // [6, 1, 2, 8, 7, 5, 3, 4, 9]
		{0, 0, 0, 6, 0, 1, 0, 7, 8}, // [4, 9, 3, 6, 2, 1, 5, 7, 8]
		{0, 0, 7, 0, 4, 0, 2, 6, 0}, // [8, 5, 7, 9, 4, 3, 2, 6, 1]
		{0, 0, 1, 0, 5, 0, 9, 3, 0}, // [2, 6, 1, 7, 5, 8, 9, 3, 4]
		{9, 0, 4, 0, 6, 0, 0, 0, 5}, // [9, 3, 4, 1, 6, 2, 7, 8, 5]
		{0, 7, 0, 3, 0, 0, 0, 1, 2}, // [5, 7, 8, 3, 9, 4, 6, 1, 2]
		{1, 2, 0, 0, 0, 7, 4, 0, 0}, // [1, 2, 6, 5, 8, 7, 4, 9, 3]
		{0, 4, 9, 2, 0, 6, 0, 0, 7}, // [3, 4, 9, 2, 1, 6, 8, 5, 7]
	}

	fmt.Println(recursion.SolveSudoku(board))
}
