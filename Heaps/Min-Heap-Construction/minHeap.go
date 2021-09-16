package minheap

import (
	"math"
)

type MinHeap struct {
	Heap []int
}

// Big O: O(n) time | O(1) space
func New(array []int) *MinHeap {
	firstParentIdx := int(math.Floor(float64(len(array)-2) / 2))
	// encounter a weird case here for problem [Heaps/Sorted K-Sorted Array]
	// if we pass 'Heap: array' directly below, when sorting the same array
	// in place, the heap value is accidently change too, so make a copy here
	copyArr := make([]int, len(array))
	copy(copyArr, array)
	minHeap := &MinHeap{
		Heap: copyArr,
	}
	for currentIdx := firstParentIdx; currentIdx >= 0; currentIdx-- {
		minHeap.siftDown(currentIdx, len(array)-1)
	}
	return minHeap
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	childOneIdx := 2*currentIdx + 1
	var childTwoIdx int
	for childOneIdx <= endIdx {
		if 2*currentIdx+2 <= endIdx {
			childTwoIdx = 2*currentIdx + 2
		} else {
			childTwoIdx = -1
		}
		idxToSwap := currentIdx
		if childTwoIdx != -1 && h.Heap[childTwoIdx] < h.Heap[childOneIdx] {
			idxToSwap = childTwoIdx
		} else {
			idxToSwap = childOneIdx
		}
		if h.Heap[currentIdx] > h.Heap[idxToSwap] {
			h.swap(currentIdx, idxToSwap)
			currentIdx = idxToSwap
			childOneIdx = 2*currentIdx + 1
		} else {
			break
		}
	}
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftUp(currentIdx int) {
	parentIdx := int(math.Floor((float64(currentIdx - 1)) / 2))
	for currentIdx > 0 && h.Heap[currentIdx] < h.Heap[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = int(math.Floor((float64(currentIdx - 1)) / 2))
	}
}

func (h *MinHeap) Peak() int {
	return h.Heap[0]
}

// O(log(n)) time | O(1) space
func (h *MinHeap) Remove() int {
	h.swap(0, len(h.Heap)-1)
	valueToRemove := h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]
	h.siftDown(0, len(h.Heap)-1)
	// fmt.Printf("value to remove: %v, heap: %v\n", valueToRemove, h.Heap)
	return valueToRemove
}

// O(log(n)) time | O(1) space
func (h *MinHeap) Insert(value int) {
	h.Heap = append(h.Heap, value)
	h.siftUp(len(h.Heap) - 1)
}

func (h *MinHeap) swap(i, j int) {
	(*h).Heap[i], (*h).Heap[j] = (*h).Heap[j], (*h).Heap[i]
}
