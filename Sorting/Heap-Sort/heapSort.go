package sort

import (
	"math"
)

// Big O: O(n*log(n)) time | O(1) space
// total n iterations, for each iteration there's one swap O(1) time and one
// siftDown O(log(n)) time, therefore O(n*log(n)) time
func HeapSort(array []int) {
	buildMaxHeap(array)
	// imagine the original array is divided into two parts:
	// [ [a max heap with root value as current maximum value] | [sorted array, initially empty] ]
	// keep iterating from the back of the array, swap the root value of the heap with the
	// iterating index at the end, then maintain the max heap feature by sifting down
	// stop when the heap slice reduces to length 1
	for endIdx := len(array) - 1; endIdx > 0; endIdx-- {
		swap(array, 0, endIdx)
		siftDown(array, 0, endIdx-1)
	}
}

// Big O: O(n) time | O(1) space
func buildMaxHeap(array []int) {
	firstParentIdx := int(math.Floor(float64(len(array)-2) / 2))
	for currentIdx := firstParentIdx; currentIdx >= 0; currentIdx-- {
		siftDown(array, currentIdx, len(array)-1)
	}
}

// O(log(n)) time | O(1) space
func siftDown(heap []int, currentIdx, endIdx int) {
	childOneIdx := 2*currentIdx + 1
	var childTwoIdx int
	for childOneIdx <= endIdx {
		if 2*currentIdx+2 <= endIdx {
			childTwoIdx = 2*currentIdx + 2
		} else {
			childTwoIdx = -1
		}
		idxToSwap := currentIdx
		if childTwoIdx != -1 && heap[childTwoIdx] > heap[childOneIdx] {
			idxToSwap = childTwoIdx
		} else {
			idxToSwap = childOneIdx
		}
		if heap[currentIdx] < heap[idxToSwap] {
			swap(heap, currentIdx, idxToSwap)
			currentIdx = idxToSwap
			childOneIdx = 2*currentIdx + 1
		} else {
			break
		}
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
