package main

import (
	"fmt"
	"math"
)

func main() {
	srcArray := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	fmt.Println(kadanesAlgorithm(srcArray))
}

// Big O: O(n) time | O(1) space
// maxEndingHere represents the sum of subarray from index 0 to every iteration index i, i.e. for subarray array[0:i],
// max sum ending here(for this subarray) is the larger value between a) previous maxEndingHere (value at previous index) plus current value,
// or b) value at current index alone
func kadanesAlgorithm(array []int) int {
	maxEndingHere, maxSoFar := array[0], array[0]
	for i := 1; i < len(array); i++ {
		maxEndingHere = int(math.Max(float64(maxEndingHere+array[i]), float64(array[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}
	return maxSoFar
}
