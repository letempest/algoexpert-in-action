package main

import (
	"fmt"
)

func main() {
	srcArray := make([][]int, 5)
	srcArray[0] = []int{1, 3, 4, 10, 11}
	srcArray[1] = []int{2, 5, 9, 12, 19}
	srcArray[2] = []int{6, 8, 13, 18, 20}
	srcArray[3] = []int{7, 14, 17, 21, 24}
	srcArray[4] = []int{15, 16, 22, 23, 25}

	fmt.Println(zigzagTraverse(srcArray))
}

// // My solution
// // Big O: O(n) time | O(n) space
// func zigzagTraverse(array [][]int) []int {
// 	prevDirection := -1
// 	var nextDirection int
// 	result := make([]int, 0)
// 	for i, j := 0, 0; i < len(array) && j < len(array[0]); {
// 		if i == 0 && j == 0 {
// 			nextDirection = getNextDirection(i, j, len(array)-1, len(array[i])-1, -1)
// 		} else {
// 			nextDirection = getNextDirection(i, j, len(array)-1, len(array[i])-1, prevDirection)
// 		}
// 		result = append(result, array[i][j])
// 		switch nextDirection {
// 		case 1:
// 			i += 1
// 		case 2:
// 			j += 1
// 		case 3:
// 			i -= 1
// 			j += 1
// 		case 4:
// 			i += 1
// 			j -= 1
// 		}
// 		prevDirection = nextDirection
// 	}
// 	return result
// }

// // for every element, check where to go next: 1 = go down; 2 = go right; 3 = go 45deg up; 4 = go -135deg down
// func getNextDirection(i, j, x, y, prevDirection int) int {
// 	var nextDirection int
// 	switch prevDirection {
// 	case -1:
// 		nextDirection = 1
// 	case 1:
// 		if j == 0 {
// 			nextDirection = 3
// 		} else if i == x {
// 			nextDirection = 2
// 		} else {
// 			nextDirection = 4
// 		}
// 	case 2:
// 		if i == 0 {
// 			nextDirection = 4
// 		} else {
// 			nextDirection = 3
// 		}
// 	case 3:
// 		if i-1 >= 0 && j+1 <= y {
// 			nextDirection = 3
// 		} else if i == 0 && j+1 <= y {
// 			nextDirection = 2
// 		} else {
// 			nextDirection = 1
// 		}
// 	case 4:
// 		if i+1 <= x && j-1 >= 0 {
// 			nextDirection = 4
// 		} else if i == x {
// 			nextDirection = 2
// 		} else {
// 			nextDirection = 1
// 		}
// 	default:
// 		panic(fmt.Sprintf("at row: %v, col: %v, yield a non 1-4 value direction.\n", i, j))
// 	}

// 	return nextDirection
// }

// Expert Solution
// Big O: O(n) time | O(n) space
func zigzagTraverse(array [][]int) []int {
	height := len(array) - 1
	width := len(array[0]) - 1
	result := make([]int, 0)
	goingDown := true

	for row, col := 0, 0; !isOutOfBounds(row, col, height, width); {
		result = append(result, array[row][col])
		if goingDown {
			if col == 0 || row == height {
				goingDown = false
				if row == height {
					col += 1
				} else {
					row += 1
				}
			} else {
				row += 1
				col -= 1
			}
		} else {
			if row == 0 || col == width {
				goingDown = true
				if col == width {
					row += 1
				} else {
					col += 1
				}
			} else {
				row -= 1
				col += 1
			}
		}
	}
	return result
}

func isOutOfBounds(row, col, height, width int) bool {
	return row < 0 || row > height || col < 0 || col > width
}
