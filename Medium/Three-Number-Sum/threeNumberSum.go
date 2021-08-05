package main

import (
	"fmt"
	"sort"
)

// Give an unsorted array, return all triplet arrays with three elements that sum up to equal the targetSum number.
// Big O: O(n^2) Time | O(n) Space
func main() {
	var arr = []int{12, 3, 1, 2, -6, 5, -8, 6}
	fmt.Println(threeNumberSum(arr, 0))
}

func threeNumberSum(arr []int, targetSum int) [][]int {
	sort.Ints(arr)

	var left, right int
	triplets := [][]int{}

	for i := 0; i < len(arr)-2; i++ {
		left = i + 1
		right = len(arr) - 1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum < targetSum {
				left += 1
			} else if sum > targetSum {
				right -= 1
			} else {
				triplets = append(triplets, []int{arr[i], arr[left], arr[right]})
				left++
				right--
			}
		}
	}

	return triplets
}
