// Given two array of numbers, find out whether they can construct a same
// bst or not
// Conceptual overview: recursive problem, for each recursive call, we know
// the current root value, filter out two subarrays for left/right subtree
// for both trees, and call the same function on the four subarrays, until
// we reach the base case where the arrays (in other word, the left/right)
// subtree to be compared are either empty or of length 1, for which we only
// need to compare the first elements.

package bst

import "math"

// Solution 1
// Big O: O(n^2) time | O(n^2) space
// func SameBSTs(arrayOne, arrayTwo []int) bool {
// 	if len(arrayOne) != len(arrayTwo) {
// 		return false
// 	}
// 	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
// 		return true
// 	}
// 	if arrayOne[0] != arrayTwo[0] {
// 		return false
// 	}

// 	currentRoot := arrayOne[0]

// 	leftSubtreeOne := make([]int, 0, len(arrayOne))
// 	rightSubtreeOne := make([]int, 0, len(arrayOne))
// 	leftSubtreeTwo := make([]int, 0, len(arrayTwo))
// 	rightSubtreeTwo := make([]int, 0, len(arrayTwo))

// 	for i := 1; i < len(arrayOne); i++ {
// 		if arrayOne[i] < currentRoot {
// 			leftSubtreeOne = append(leftSubtreeOne, arrayOne[i])
// 		} else {
// 			rightSubtreeOne = append(rightSubtreeOne, arrayOne[i])
// 		}
// 		if arrayTwo[i] < currentRoot {
// 			leftSubtreeTwo = append(leftSubtreeTwo, arrayTwo[i])
// 		} else {
// 			rightSubtreeTwo = append(rightSubtreeTwo, arrayTwo[i])
// 		}
// 	}

// 	return SameBSTs(leftSubtreeOne, leftSubtreeTwo) && SameBSTs(rightSubtreeOne, rightSubtreeTwo)
// }

// =================================================================
// Improved Solution
// Big O: O(n^2) time, conceptually the same as solution above
// improved space complexity: O(d) space, where d is the depth of the bst, at most d recursive calls
// on the callstack; always keep track of the index and value bound of the left/right subtree
// for the next recursive call, so no additional space for left/right subtree value slices
func SameBSTs(arrayOne, arrayTwo []int) bool {
	return areSameBSTs(arrayOne, arrayTwo, 0, 0, math.MinInt64, math.MaxInt64)
}

func areSameBSTs(arrayOne, arrayTwo []int, rootIdxOne, rootIdxTwo int, minVal, maxVal int) bool {
	if rootIdxOne == -1 || rootIdxTwo == -1 {
		return rootIdxOne == rootIdxTwo
	}

	if arrayOne[rootIdxOne] != arrayTwo[rootIdxTwo] {
		return false
	}

	leftRootIdxOne := getIdxOfFirstSmaller(arrayOne, rootIdxOne, minVal)
	leftRootIdxTwo := getIdxOfFirstSmaller(arrayTwo, rootIdxTwo, minVal)
	rightRootIdxOne := getIdxOfFirstBiggerOrEqual(arrayOne, rootIdxOne, maxVal)
	rightRootIdxTwo := getIdxOfFirstBiggerOrEqual(arrayTwo, rootIdxTwo, maxVal)

	currentRootVal := arrayOne[rootIdxOne]
	leftAreSame := areSameBSTs(arrayOne, arrayTwo, leftRootIdxOne, leftRootIdxTwo, minVal, currentRootVal)
	rightAreSame := areSameBSTs(arrayOne, arrayTwo, rightRootIdxOne, rightRootIdxTwo, currentRootVal, maxVal)

	return leftAreSame && rightAreSame
}

func getIdxOfFirstSmaller(array []int, startingIdx int, minVal int) int {
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] < array[startingIdx] && array[i] >= minVal {
			return i
		}
	}
	return -1
}

func getIdxOfFirstBiggerOrEqual(array []int, startingIdx int, maxVal int) int {
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] >= array[startingIdx] && array[i] < maxVal {
			return i
		}
	}
	return -1
}
