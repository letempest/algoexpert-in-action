package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%#v\n", searchForRange([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}, 45))
}

// Solution 1: Recursive Approach
// Big O: O(log(n)) time | O(log(n)) space, at most log(n) recursive calls on the callstack
// func searchForRange(array []int, target int) [2]int {
// 	finalRange := [2]int{-1, -1}
// 	alteredBinarySearch(array, target, 0, len(array)-1, &finalRange, true)
// 	alteredBinarySearch(array, target, 0, len(array)-1, &finalRange, false)
// 	return finalRange
// }

// func alteredBinarySearch(array []int, target, leftIdx, rightIdx int, finalRange *[2]int, goLeft bool) {
// 	if leftIdx > rightIdx {
// 		return
// 	}
// 	midIdx := int(math.Floor(float64(leftIdx+rightIdx) / 2))
// 	if array[midIdx] < target {
// 		alteredBinarySearch(array, target, midIdx+1, rightIdx, finalRange, goLeft)
// 	} else if array[midIdx] > target {
// 		alteredBinarySearch(array, target, leftIdx, midIdx-1, finalRange, goLeft)
// 	} else {
// 		if goLeft {
// 			if midIdx == 0 || array[midIdx-1] != target {
// 				finalRange[0] = midIdx
// 			} else {
// 				alteredBinarySearch(array, target, leftIdx, midIdx-1, finalRange, goLeft)
// 			}
// 		} else {
// 			if midIdx == len(array)-1 || array[midIdx+1] != target {
// 				finalRange[1] = midIdx
// 			} else {
// 				alteredBinarySearch(array, target, midIdx+1, rightIdx, finalRange, goLeft)
// 			}
// 		}
// 	}
// }

// ==========================================================================
// Solution 2: Iterative Approach
// O(log(n)) time | O(1) space
func searchForRange(array []int, target int) [2]int {
	finalRange := [2]int{-1, -1}
	alteredBinarySearch(array, target, 0, len(array)-1, &finalRange, true)  // one recursive call to keep searching on the left side
	alteredBinarySearch(array, target, 0, len(array)-1, &finalRange, false) // similarly, update the rightmost range by searching on the right side
	return finalRange
}

func alteredBinarySearch(array []int, target, leftIdx, rightIdx int, finalRange *[2]int, goLeft bool) {
	var midIdx int
	for leftIdx <= rightIdx {
		midIdx = int(math.Floor(float64(leftIdx+rightIdx) / 2))
		if array[midIdx] < target {
			leftIdx = midIdx + 1
		} else if array[midIdx] > target {
			rightIdx = midIdx - 1
		} else {
			if goLeft {
				if midIdx == 0 || array[midIdx-1] != target {
					finalRange[0] = midIdx
					return
				} else {
					rightIdx = midIdx - 1
				}
			} else {
				if midIdx == len(array)-1 || array[midIdx+1] != target {
					finalRange[1] = midIdx
					return
				} else {
					leftIdx = midIdx + 1
				}
			}
		}
	}
}
