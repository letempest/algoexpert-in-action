package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	// should return index 3
	fmt.Println(binarySearch(input, 33))
	// no match, should return -1
	fmt.Println(binarySearch(input, 39))
}

// Iterative Solution
// Big O: O(log(n)) time | O(1) space
func binarySearch(array []int, target int) int {
	leftIdx := 0
	rightIdx := len(array) - 1

	for midIdx := 0; leftIdx <= rightIdx; {
		midIdx = int(math.Floor((float64(leftIdx) + float64(rightIdx)) / 2))
		if array[midIdx] < target {
			leftIdx = midIdx + 1
		} else if array[midIdx] > target {
			rightIdx = midIdx - 1
		} else {
			return midIdx
		}
	}
	return -1
}

// Recursive Solution
// Big O: O(log(n)) time | O(log(n)) space, at most log(n) recursive calls on the callstack
// func binarySearch(array []int, target int) int {
// 	return binarySearchRecurHelper(array, target, 0, len(array)-1)
// }

// func binarySearchRecurHelper(array []int, target, leftIdx, rightIdx int) int {
// 	// exit condition
// 	if leftIdx > rightIdx {
// 		// intentionally return a negative index to signal Not Found
// 		return -1
// 	}

// 	midIdx := int(math.Floor(((float64(leftIdx) + float64(rightIdx)) / 2)))
// 	if array[midIdx] < target {
// 		return binarySearchRecurHelper(array, target, midIdx+1, rightIdx)
// 	} else if array[midIdx] > target {
// 		return binarySearchRecurHelper(array, target, leftIdx, midIdx-1)
// 	} else {
// 		// find matching
// 		return midIdx
// 	}
// }
