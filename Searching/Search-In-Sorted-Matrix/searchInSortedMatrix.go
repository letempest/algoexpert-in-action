// Given a sorted matrix in which each row and column is SORTED and a target value, return the row, coloumn index of matching result

package main

import "fmt"

var matrix [][]int

func init() {
	matrix = make([][]int, 5)
	matrix[0] = []int{1, 4, 7, 12, 15, 1000}
	matrix[1] = []int{2, 5, 19, 31, 32, 1001}
	matrix[2] = []int{3, 8, 24, 33, 35, 1002}
	matrix[3] = []int{40, 41, 42, 44, 45, 1003}
	matrix[4] = []int{99, 100, 103, 106, 128, 1004}
}

func main() {
	target := 99
	row, col := searchInSortedMatrix(matrix, target)
	fmt.Printf("The row and column index of value %v in matrix is: %v, %v respectively\n", target, row, col)
}

// Big O: O(n + m) time, where n, m are number of rows and cols of the matrix; O(1) time
func searchInSortedMatrix(matrix [][]int, target int) (int, int) {
	row, col := len(matrix), len(matrix[0])
	for i, j := 0, col-1; i < row && j >= 0; {
		currentVal := matrix[i][j]
		if currentVal > target {
			j -= 1
		} else if currentVal < target {
			i += 1
		} else {
			return i, j
		}
	}
	return -1, -1
}
