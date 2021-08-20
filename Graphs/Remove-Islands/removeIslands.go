// Island: node with value equal to 1 and NOT connected to border (means its adjacent border node value could not be 1 either)
// and replace such island's value with 0

package main

import "fmt"

var matrix [][]int

func init() {
	matrix = make([][]int, 6)
	matrix[0] = []int{1, 0, 0, 0, 0, 0} // 1, 0, 0, 0, 0, 0		r
	matrix[1] = []int{0, 1, 0, 1, 1, 1} // 0, 0, 0, 1, 1, 1		e
	matrix[2] = []int{0, 0, 1, 0, 1, 0} // 0, 0, 0, 0, 1, 0		s
	matrix[3] = []int{1, 1, 0, 0, 1, 0} // 1, 1, 0, 0, 1, 0		u
	matrix[4] = []int{1, 0, 1, 1, 0, 0} // 1, 0, 0, 0, 0, 0		l
	matrix[5] = []int{1, 0, 0, 0, 0, 1} // 1, 0, 0, 0, 0, 1		t
}

func main() {
	fmt.Println(removeIslands((matrix)))
}

// Solution 1: auxiliary matrix to mark a corresponding node as (node with value 1 and is connected to border)
// Big O: O(wh) time | O(wh) space; where w, h represent width and length of matrix respectively
// Space complexity explantion: 1) same size auxiliary matrix;
// 2) if matrix is filled with 1, the stack will be at length w*h, since every node will be pushed into it
// func removeIslands(matrix [][]int) [][]int {
// 	// Init an identical onesConnectedToBoarder matrix
// 	onesConnectedToBoarder := make([][]bool, len(matrix))
// 	for i := range matrix {
// 		onesConnectedToBoarder[i] = make([]bool, len(matrix[i]))
// 		for j := range matrix[i] {
// 			onesConnectedToBoarder[i][j] = false
// 		}
// 	}

// 	for row := range matrix {
// 		for col := range matrix[row] {
// 			rowIsBorder := row == 0 || row == len(matrix)-1
// 			colIsBorder := col == 0 || col == len(matrix[row])-1
// 			isBorder := rowIsBorder || colIsBorder
// 			// Not at boarder, don't do anything
// 			if !isBorder {
// 				continue
// 			}

// 			if matrix[row][col] != 1 {
// 				continue
// 			}

// 			// only call this func on all boundry nodes with value equal to 1
// 			findOnesConnectedToBoarder(matrix, row, col, onesConnectedToBoarder)
// 		}
// 	}

// 	for i := 1; i < len(matrix)-1; i++ {
// 		for j := 1; j < len(matrix[i])-1; j++ {
// 			if !onesConnectedToBoarder[i][j] {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// 	return matrix
// }

// func findOnesConnectedToBoarder(matrix [][]int, row, col int, onesConnectedToBoarder [][]bool) {
// 	stack := [][]int{{row, col}}
// 	for len(stack) > 0 {
// 		currentNode := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		i := currentNode[0]
// 		j := currentNode[1]
// 		alreadyVisited := onesConnectedToBoarder[i][j]
// 		if alreadyVisited {
// 			continue
// 		}

// 		onesConnectedToBoarder[i][j] = true

// 		// get the neighbors of current node, if the value of neighbor equals to 1, then push it to the stack
// 		neighbors := getNeighbors(matrix, i, j)
// 		for _, indexes := range neighbors {
// 			row, col := indexes[0], indexes[1]
// 			if matrix[row][col] == 1 {
// 				// ensure that every node push to the stack has a value equal to 1
// 				stack = append(stack, indexes)
// 			}
// 		}
// 	}
// }

// func getNeighbors(matrix [][]int, row, col int) [][]int {
// 	var neighbors [][]int
// 	// UP
// 	if row > 0 {
// 		neighbors = append(neighbors, []int{row - 1, col})
// 	}
// 	// DOWN
// 	if row < len(matrix)-1 {
// 		neighbors = append(neighbors, []int{row + 1, col})
// 	}
// 	// LEFT
// 	if col > 0 {
// 		neighbors = append(neighbors, []int{row, col - 1})
// 	}
// 	// RIGHT
// 	if col < len(matrix[row])-1 {
// 		neighbors = append(neighbors, []int{row, col + 1})
// 	}
// 	return neighbors
// }

// =====================================================================
// Solution 2: mutating matrix values in place, no auxiliary matrix, slightly better space complexity (though at the same scale)
// Big O: O(wh) time | O(wh) space; where w, h represent width and length of matrix respectively
func removeIslands(matrix [][]int) [][]int {
	for row := range matrix {
		for col := range matrix[row] {
			rowIsBorder := row == 0 || row == len(matrix)-1
			colIsBorder := col == 0 || col == len(matrix[row])-1
			isBorder := rowIsBorder || colIsBorder
			// Not at boarder, don't do anything
			if !isBorder {
				continue
			}

			if matrix[row][col] != 1 {
				continue
			}

			// only call this func on all boundry nodes with value equal to 1
			changeOnesConnectedToBoarderToTwos(matrix, row, col)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = 0
			}
			if matrix[i][j] == 2 {
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}

func changeOnesConnectedToBoarderToTwos(matrix [][]int, row, col int) {
	stack := [][]int{{row, col}}
	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i := currentNode[0]
		j := currentNode[1]

		matrix[i][j] = 2

		// get the neighbors of current node, if the value of neighbor equals to 1, then push it to the stack
		neighbors := getNeighbors(matrix, i, j)
		for _, indexes := range neighbors {
			row, col := indexes[0], indexes[1]
			if matrix[row][col] == 1 {
				// ensure that every node push to the stack has a value equal to 1
				stack = append(stack, indexes)
			}
		}
	}
}

func getNeighbors(matrix [][]int, row, col int) [][]int {
	var neighbors [][]int
	// UP
	if row > 0 {
		neighbors = append(neighbors, []int{row - 1, col})
	}
	// DOWN
	if row < len(matrix)-1 {
		neighbors = append(neighbors, []int{row + 1, col})
	}
	// LEFT
	if col > 0 {
		neighbors = append(neighbors, []int{row, col - 1})
	}
	// RIGHT
	if col < len(matrix[row])-1 {
		neighbors = append(neighbors, []int{row, col + 1})
	}
	return neighbors
}
