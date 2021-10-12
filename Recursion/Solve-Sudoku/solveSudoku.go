package recursion

import (
	"math"
)

// Big O: O(1) time | O(1) space
// for fix-sized 9 by 9 sudoku board, there's no variable influencing the execution time of the
// algorithm, it's like a O(constant * 1) algorithm, the constant is large, yet still a constant;
// for space complexity, the course video suggests that values are filled in place so O(1),
// but actually the recursive function takes extra space on the callstack,
// yet for the same reason above (fix-sized input), the possible amount of recursive calls is
// also bounded (by a large value, yet still constant), therefore O(1) space
func SolveSudoku(board [9][9]int) [9][9]int {
	solvePartialSudoku(0, 0, &board)
	return board
}

func solvePartialSudoku(row, col int, board *[9][9]int) bool {
	currentRow := row
	currentCol := col
	// base case
	if currentCol == len(board[row]) {
		currentRow += 1
		currentCol = 0
		// if we reach beyond the last row and did not backtrack (return false),
		// then we have succeeded in solving the sudoku
		if currentRow == len(board) {
			return true
		}
	}

	// try to fill current position with 1-9
	if board[currentRow][currentCol] == 0 {
		return tryDigitsAtPosition(currentRow, currentCol, board)
	}

	// if we reach here, current position is already filled,
	// try next slot
	return solvePartialSudoku(currentRow, currentCol+1, board)
}

func tryDigitsAtPosition(row, col int, board *[9][9]int) bool {
	for digit := 1; digit < 10; digit++ {
		if isValidAtPosition(digit, row, col, board) {
			board[row][col] = digit
			if solvePartialSudoku(row, col+1, board) {
				return true
			}
		}
	}

	// if we make out of the for loop but did not return true,
	// then current position cannot be filled with 1-9,
	// which indicates slot(s) we filled before is incorrect,
	// return to previous recursive calls to try other digits
	board[row][col] = 0
	return false
}

func isValidAtPosition(value, row, col int, board *[9][9]int) bool {
	rowIsValid, colIsValid := true, true
	for _, num := range board[row] {
		if num == value {
			rowIsValid = false
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][col] == value {
			colIsValid = false
		}
	}
	if !rowIsValid || !colIsValid {
		return false
	}
	// check for sub-grid
	subGridRowStart := int(math.Floor(float64(row) / 3))
	subGridColStart := int(math.Floor(float64(col) / 3))
	// look through the nine slot in the same sub-grid,
	// if same value found, then apparently invalid
	for rowIdx := 0; rowIdx < 3; rowIdx++ {
		rowToCheck := subGridRowStart*3 + rowIdx
		for colIdx := 0; colIdx < 3; colIdx++ {
			colToCheck := subGridColStart*3 + colIdx
			existingValue := board[rowToCheck][colToCheck]
			if existingValue == value {
				return false
			}
		}
	}

	return true
}
