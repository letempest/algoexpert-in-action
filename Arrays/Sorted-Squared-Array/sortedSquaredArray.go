package array

import "math"

// Big O: O(n) time | O(n) space
func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))
	smallerValueIdx, largerValueIdx := 0, len(array)-1

	for i := len(array) - 1; i >= 0; i-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]
		if math.Abs(float64(smallerValue)) < math.Abs(float64(largerValue)) {
			sortedSquares[i] = int(math.Pow(float64(largerValue), 2))
			largerValueIdx -= 1
		} else {
			sortedSquares[i] = int(math.Pow(float64(smallerValue), 2))
			smallerValueIdx += 1
		}
	}

	return sortedSquares
}
