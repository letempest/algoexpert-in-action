// Given a rectangular graph/grid, starting from top-left node, and ending at bottom-right node, calculate number of paths could be taken
// limitation: can only go right/down for each step; (width * height >= 2) => unique start/end node

package main

import "fmt"

func main() {
	fmt.Println(numberOfWaysToTraverseGraph(4, 3))
}

// Solution 1, Recursive - Big O: O(2^(n + m)) time where n, m is the dimension of the trid; | O(n+m) space, n+m recursive calls on the callstack
// func numberOfWaysToTraverseGraph(width, height int) int {
// 	// Base case, at the border of rectangle grid
// 	if width == 1 || height == 1 {
// 		return 1
// 	}

// 	return numberOfWaysToTraverseGraph(width-1, height) + numberOfWaysToTraverseGraph(width, height-1)
// }

// =========================================================================

// Solution 2, Dynamic Programming
// Big O: O(n * m) time where n, m is the dimension of the trid, because we're looping through all the nodes for once
// O(n * m) space, because there's an auxiliary matrix for intermediate result storage
func numberOfWaysToTraverseGraph(width, height int) int {
	// Initialize a two dimensional array with identical size to store intermediate result
	numberOfWaysForEachNode := make([][]int, height)
	for i := range numberOfWaysForEachNode {
		numberOfWaysForEachNode[i] = make([]int, width)
	}

	for widthIdx := 0; widthIdx < width; widthIdx++ {
		for heightIdx := 0; heightIdx < height; heightIdx++ {
			if widthIdx == 0 || heightIdx == 0 {
				numberOfWaysForEachNode[heightIdx][widthIdx] = 1
			} else {
				waysLeft := numberOfWaysForEachNode[heightIdx][widthIdx-1]
				waysUp := numberOfWaysForEachNode[heightIdx-1][widthIdx]
				numberOfWaysForEachNode[heightIdx][widthIdx] = waysLeft + waysUp
			}
		}
	}
	return numberOfWaysForEachNode[height-1][width-1]
}

// =========================================================================

// Solution 3: leveraging math deduction, given the fact that we can only go right/down, result set would be somewhat related to
// permutation of [right,right,right..., down, down, ..., down], where there's width amount of right element, and height amount of down element
//                       \ |  /               \   |   /
//                       width                  height
// and permutation of n yields a result of n!
// O(n + m) time | O(1) space
// func numberOfWaysToTraverseGraph(width, height int) int {
// 	xDistanceToCorner := width - 1
// 	yDistanceToCorner := height - 1
// 	return factorial(xDistanceToCorner+yDistanceToCorner) / (factorial(xDistanceToCorner) * factorial(yDistanceToCorner))
// }

// func factorial(num int) int {
// 	result := 1
// 	for n := 2; n < num+1; n++ {
// 		result *= n
// 	}
// 	return result
// }
