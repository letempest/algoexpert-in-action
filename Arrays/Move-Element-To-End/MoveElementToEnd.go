// Given an array and a target value, move all elements with equal value to the end of the src array (in place!)
package main

import "fmt"

func main() {
	result := moveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2)
	fmt.Println(result)
	fmt.Println(moveElementToEnd([]int{4, 3, 7, 2, 5, 3, 0, 3, 9, -1, 3}, 3))
}

// Big O: O(n) time | O(1) space
func moveElementToEnd(array []int, target int) []int {
	for i, j := 0, len(array)-1; i < j; {
		if array[j] == target {
			j -= 1
		} else {
			if array[i] == target {
				array[i], array[j] = array[j], array[i]
				j -= 1
			}
			i += 1
		}
		// fmt.Printf("slices so far: %v, i: %v, j: %v\n", array, i, j)
	}
	return array
}
