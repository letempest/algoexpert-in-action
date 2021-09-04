package main

import (
	"fmt"
)

func main() {
	srcArray := []int{8, 12, 2, 3, 15, 5, 7}
	sum, subSequence := maxSumIncreasingSubsequence(srcArray)
	fmt.Println(sum, subSequence)
}

// Big O: O(n^2) time | O(n) space
func maxSumIncreasingSubsequence(array []int) (int, []int) {
	sums := make([]int, len(array))
	sequences := make([]int, len(array))
	sums[0] = array[0]
	for i := 0; i < len(array); i++ {
		sequences[i] = -1
	}
	maxSumIdx := 0

	for i, currentSum := range array {
		for j := 0; j < i; j++ {
			otherNum := array[j]
			if otherNum < currentSum && sums[j]+currentSum >= sums[i] {
				sums[i] = sums[j] + currentSum
				sequences[i] = j
			}
		}
		if sums[i] >= sums[maxSumIdx] {
			maxSumIdx = i
		}
	}

	return sums[maxSumIdx], buildSequence(array, sequences, maxSumIdx)
}

func buildSequence(array, sequences []int, currentIdx int) []int {
	sequence := []int{array[currentIdx]}
	for sequences[currentIdx] != -1 {
		sequence = append(sequence, array[sequences[currentIdx]])
		currentIdx = sequences[currentIdx]
	}

	for i, j := 0, len(sequence)-1; i < j; {
		sequence[i], sequence[j] = sequence[j], sequence[i]
		i += 1
		j -= 1
	}

	return sequence
}
