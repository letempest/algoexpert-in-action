package main

import "fmt"

func main() {
	matrix := make([][]int, 4)
	matrix[0] = []int{1, 2, 3, 4}
	matrix[1] = []int{12, 13, 14, 5}
	matrix[2] = []int{11, 16, 15, 6}
	matrix[3] = []int{10, 9, 8, 7}

	matrix2 := make([][]int, 5)
	matrix2[0] = []int{1, 2, 3, 4, 5, 6}
	matrix2[1] = []int{18, 19, 20, 21, 22, 7}
	matrix2[2] = []int{17, 28, 29, 30, 23, 8}
	matrix2[3] = []int{16, 27, 26, 25, 24, 9}
	matrix2[4] = []int{15, 14, 13, 12, 11, 10}

	fmt.Println("Iterative solution:")
	fmt.Printf("\t%v\n", spiralTraverse(matrix))
	fmt.Printf("\t%v\n\n", spiralTraverse(matrix2))

	fmt.Println("Recursive solution:")
	fmt.Printf("\t%v\n", spiralTraverseRecur(matrix))
	fmt.Printf("\t%v\n", spiralTraverseRecur(matrix2))
}

// My solution: iterative approach
// Big O: O(N) time, where N is total number of elements in the matrix | O(N) space
func spiralTraverse(matrix [][]int) []int {
	result := []int{}

	startRow, endRow := 0, len(matrix)-1
	startCol, endCol := 0, len(matrix[0])-1
	for startRow <= endRow && startCol <= endCol {
		// One perimeter cycle
		i, j := startRow, startCol
		// traversing starting row
		for i == startRow && j < endCol {
			result = append(result, matrix[i][j])
			j += 1
		}
		// traversing ending column
		for j == endCol && i < endRow {
			result = append(result, matrix[i][j])
			i += 1
		}
		// traversing ending row
		for i == endRow && j > startCol {
			result = append(result, matrix[i][j])
			j -= 1
		}
		// traversing starting column
		for j == startCol && i > startRow {
			result = append(result, matrix[i][j])
			i -= 1
		}
		startRow += 1
		endRow -= 1
		startCol += 1
		endCol -= 1
		i = startRow
		j = startCol
	}

	return result
}

// Recursive solution
// Big O: O(N) time | O(N) space, recursive calls on callstack is layers of perimeter, roughly min(n, m)/2, says min(n, m), still definitely smaller than O(N)
func spiralTraverseRecur(matrix [][]int) []int {
	result := &[]int{}
	spiralTraverseRecurHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, result)
	return *result
}

func spiralTraverseRecurHelper(matrix [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	// base case
	if startRow > endRow || startCol > endCol {
		return
	}

	// One perimeter cycle
	i, j := startRow, startCol
	for i == startRow && j < endCol {
		*result = append(*result, matrix[i][j])
		j += 1
	}
	for j == endCol && i < endRow {
		*result = append(*result, matrix[i][j])
		i += 1
	}
	for i == endRow && j > startCol {
		*result = append(*result, matrix[i][j])
		j -= 1
	}
	for j == startCol && i > startRow {
		*result = append(*result, matrix[i][j])
		i -= 1
	}

	spiralTraverseRecurHelper(matrix, startRow+1, endRow-1, startCol+1, endCol-1, result)
}
