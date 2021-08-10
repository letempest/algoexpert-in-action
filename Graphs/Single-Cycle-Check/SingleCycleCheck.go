package main

import "fmt"

func main() {
	fmt.Println(hasSingleCycle([]int{2, 3, 1, -4, -4, 2}))
	fmt.Println(hasSingleCycle([]int{1, -1, -1, 1}))
}

// Big O: O(n) time | O(1) spac
func hasSingleCycle(array []int) bool {
	n := len(array)
	const STARTING_IDX = 0
	currentIdx := STARTING_IDX
	for i := 0; i < n; i++ {
		if i != 0 && currentIdx == STARTING_IDX {
			return false
		}
		// calculate the new index, could be out of bound
		currentIdx += array[currentIdx]
		// take modulo and map back into 0 - (n-1) index scope
		currentIdx = currentIdx % n
		if currentIdx < 0 {
			currentIdx += n
		}
	}
	return currentIdx == STARTING_IDX
}
