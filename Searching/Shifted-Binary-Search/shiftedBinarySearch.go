// find the target value in a somewhat sorted array, values in the array could be shifted in one direction to some amount
// could still eliminate one half of input values on each step similar to binary search

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shiftedBinarySearch([]int{45, 61, 71, 72, 73, 0, 1, 21, 33, 45}, 33))
	fmt.Println(shiftedBinarySearch([]int{45, 61, 71, 72, 73, 0, 1, 21, 33, 45}, 99))
}

// Iterative solution
// Big O: O(log(n)) time | O(1) space
func shiftedBinarySearch(array []int, target int) int {
	for leftIdx, midIdx, rightIdx := 0, 0, len(array)-1; leftIdx <= rightIdx; {
		midIdx = int(math.Floor((float64(leftIdx) + float64(rightIdx)) / 2))
		if array[midIdx] == target {
			return midIdx
		}
		if array[leftIdx] <= array[midIdx] {
			if array[leftIdx] <= target && target < array[midIdx] {
				rightIdx = midIdx - 1
			} else {
				leftIdx = midIdx + 1
			}
		} else {
			if array[midIdx] < target && target <= array[rightIdx] {
				leftIdx = midIdx + 1
			} else {
				rightIdx = midIdx - 1
			}
		}
	}
	// not found, return negative value
	return -1
}

// Recursive solution
// Big O: O(log(n)) time | O(log(n)) space, log(n) recursive calls on callstack
// func shiftedBinarySearch(array []int, target int) int {
// 	return shiftedBinarySearchRecurHelper(array, target, 0, len(array)-1)
// }

// func shiftedBinarySearchRecurHelper(array []int, target, leftIdx, rightIdx int) int {
// 	if leftIdx > rightIdx {
// 		return -1
// 	}
// 	midIdx := int(math.Floor((float64(leftIdx) + float64(rightIdx)) / 2))
// 	if array[midIdx] == target {
// 		return midIdx
// 	}

// 	if array[leftIdx] <= array[midIdx] {
// 		if array[leftIdx] <= target && target < array[midIdx] {
// 			return shiftedBinarySearchRecurHelper(array, target, leftIdx, midIdx-1)
// 		} else {
// 			return shiftedBinarySearchRecurHelper(array, target, midIdx+1, rightIdx)
// 		}
// 	} else {
// 		if array[midIdx] < target && target <= array[rightIdx] {
// 			return shiftedBinarySearchRecurHelper(array, target, midIdx+1, rightIdx)
// 		} else {
// 			return shiftedBinarySearchRecurHelper(array, target, leftIdx, midIdx-1)
// 		}
// 	}
// }
