// Given an array, return its powerset (all possible subsets, from empty array, single element, ..., one element less, all the way up to src array itself)

package main

import "fmt"

func main() {
	srcArray := []int{1, 2, 3, 4}
	fmt.Println(powerset(srcArray))
}

// Solution 1: Iterative, O(n*2^n) time | O(n*2^n) space
// 2^n subsets in total, to create an array for each of them, with length ranging from 0 to n,
// could be an O(n/2) operation on average, O(n/2) time to generate such an array, and O(n/2) space to store it
// func powerset(slice []int) [][]int {
// 	subsets := [][]int{{}}
// 	for _, val := range slice {
// 		for _, currentSubset := range subsets {
// 			newSubset := append(currentSubset, val)
// 			subsets = append(subsets, newSubset)
// 		}
// 	}
// 	return subsets
// }

// Solution 2: Recursive, O(n*2^n) time | O(n*2^n) space, recursive calls on the callstack is no greater than n*2^n
// powerset of [1,2,3,4], i.e. P([1,2,3,4]) = adding 4 to P([1,2,3]), and P([1,2,3]) = adding 3 to P([1,2])
func powerset(slice []int, optIdx ...int) [][]int {
	var idx, ele int
	if len(optIdx) == 0 {
		idx = len(slice) - 1
	} else if optIdx[0] < 0 {
		return [][]int{{}}
	} else {
		idx = optIdx[0]
	}
	ele = slice[idx]
	subsets := powerset(slice, idx-1)
	for _, currentSubset := range subsets {
		newSubset := append(currentSubset, ele)
		subsets = append(subsets, newSubset)
	}
	return subsets
}
