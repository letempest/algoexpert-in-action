// Given a positive integer n, compute the total amount of binary tree topologies that could be constructed
// with n nodes, values do not matter, just the way these nodes are connected matter.
// e.g. for n = 3, there are 5 different topologies
//        n1      n1         n1        n1          n1
//       /       /          /  \         \           \
//      n2      n2        n2    n3        n2          n2
//     /          \                      /             \
//    n3           n3                  n3               n3

package bintreetopologies

// Solution 1, vanilla recursion approach
// Big O: Upper Bound: O(n*(2n)!/(n!*(n+1)!)) time | O(n) space, at most n calls on the callstack
// func NumberOfBinaryTreeTopologies(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	numberOfTrees := 0

// 	for leftTreeSize := 0; leftTreeSize < n; leftTreeSize++ {
// 		rightTreeSize := n - 1 - leftTreeSize
// 		numberOfLeftTrees := NumberOfBinaryTreeTopologies(leftTreeSize)
// 		numberOfRightTrees := NumberOfBinaryTreeTopologies(rightTreeSize)
// 		numberOfTrees += numberOfLeftTrees * numberOfRightTrees
// 	}
// 	return numberOfTrees
// }

// =======================================================================
// Solution 2, recursion with memoization
// O(n^2) time | O(n) space for memoized cache and at most n calls on the callstack
// func NumberOfBinaryTreeTopologies(n int) int {
// 	memoized := map[int]int{
// 		0: 1,
// 	}

// 	// TIL that map in go is reference types, no need to pass in by pointer additionally
// 	return numberOfBinaryTreeTopologiesHelper(n, memoized)
// }

// func numberOfBinaryTreeTopologiesHelper(n int, memoized map[int]int) int {
// 	if _, exists := memoized[n]; exists {
// 		return memoized[n]
// 	}

// 	numberOfTrees := 0
// 	var numberOfLeftTrees, numberOfRightTrees int

// 	for leftTreeSize := 0; leftTreeSize < n; leftTreeSize++ {
// 		rightTreeSize := n - 1 - leftTreeSize
// 		numberOfLeftTrees = numberOfBinaryTreeTopologiesHelper(leftTreeSize, memoized)
// 		numberOfRightTrees = numberOfBinaryTreeTopologiesHelper(rightTreeSize, memoized)

// 		numberOfTrees += numberOfLeftTrees * numberOfRightTrees
// 	}

// 	memoized[n] = numberOfTrees
// 	return numberOfTrees
// }

// =======================================================================
// Solution 3, iterative approach with memoization
// O(n^2) time | O(n) space
func NumberOfBinaryTreeTopologies(n int) int {
	memoized := map[int]int{
		0: 1,
	}

	for m := 1; m < n+1; m++ {
		numberOfTreesForSizeM := 0
		for leftTreeSize := 0; leftTreeSize < m; leftTreeSize++ {
			rightTreeSize := m - 1 - leftTreeSize
			numberOfLeftTrees := memoized[leftTreeSize]
			numberOfRightTrees := memoized[rightTreeSize]

			numberOfTreesForSizeM += numberOfLeftTrees * numberOfRightTrees
		}
		memoized[m] = numberOfTreesForSizeM
	}
	return memoized[n]
}
