package laptoprentals

import (
	"math"
	"sort"
)

// Solution 1
// sort the input in ascending order according to start time
// using min heap to keep track of smallest end time
// if end time <= next start time, then we can pop from the min heap
// and insert the current interval into it
// Big O: O(n * log(n)) time | O(n) space
func LaptopRentals(times [][2]int) int {
	if len(times) == 0 {
		return 0
	}
	sort.Slice(times, func(i, j int) bool { return times[i][0] < times[j][0] })
	heap := buildHeap(append([][2]int(nil), times[0]))

	for idx := 1; idx < len(times); idx++ { // looping n times
		currentInterval := times[idx]
		// for each iteration, at most 1 remove and 1 insert operation happens
		// that's 2 log(n) time operation
		if heap.peak()[1] <= currentInterval[0] {
			heap.remove()
		}
		heap.insert(currentInterval)
	}
	return len(heap.heap)
}

type minHeap struct {
	heap [][2]int
}

// Big O: O(n) time | O(1) space
func buildHeap(array [][2]int) *minHeap {
	firstParentIdx := int(math.Floor(float64(len(array)-2) / 2))
	// encounter a weird case here for problem [Heaps/Sorted K-Sorted Array]
	// if we pass 'Heap: array' directly below, when sorting the same array
	// in place, the heap value is accidently change too, so make a copy here
	copyArr := make([][2]int, len(array))
	copy(copyArr, array)
	minHeap := &minHeap{
		heap: copyArr,
	}
	for currentIdx := firstParentIdx; currentIdx >= 0; currentIdx-- {
		minHeap.siftDown(currentIdx, len(array)-1)
	}
	return minHeap
}

// O(log(n)) time | O(1) space
func (h *minHeap) siftDown(currentIdx, endIdx int) {
	childOneIdx := 2*currentIdx + 1
	var childTwoIdx int
	for childOneIdx <= endIdx {
		if 2*currentIdx+2 <= endIdx {
			childTwoIdx = 2*currentIdx + 2
		} else {
			childTwoIdx = -1
		}
		idxToSwap := currentIdx
		if childTwoIdx != -1 && h.heap[childTwoIdx][1] < h.heap[childOneIdx][1] {
			idxToSwap = childTwoIdx
		} else {
			idxToSwap = childOneIdx
		}
		if h.heap[currentIdx][1] > h.heap[idxToSwap][1] {
			h.swap(currentIdx, idxToSwap)
			currentIdx = idxToSwap
			childOneIdx = 2*currentIdx + 1
		} else {
			break
		}
	}
}

// O(log(n)) time | O(1) space
func (h *minHeap) siftUp(currentIdx int) {
	parentIdx := int(math.Floor((float64(currentIdx - 1)) / 2))
	for currentIdx > 0 && h.heap[currentIdx][1] < h.heap[parentIdx][1] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = int(math.Floor((float64(currentIdx - 1)) / 2))
	}
}

func (h *minHeap) peak() [2]int {
	return h.heap[0]
}

// O(log(n)) time | O(1) space
func (h *minHeap) remove() [2]int {
	h.swap(0, len(h.heap)-1)
	valueToRemove := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0, len(h.heap)-1)
	return valueToRemove
}

// O(log(n)) time | O(1) space
func (h *minHeap) insert(value [2]int) {
	h.heap = append(h.heap, value)
	h.siftUp(len(h.heap) - 1)
}

func (h *minHeap) swap(i, j int) {
	(*h).heap[i], (*h).heap[j] = (*h).heap[j], (*h).heap[i]
}

// ================================================================
// Solution 2
// sorting two auxiliary array startTimes and endtimes and iterate through the two
// arrays simultaneously, if end time is smaller than start time, that means
// we can free up one previously using laptop for next rental, for this case
// end time iterator is incremented and usedLaptops is decremented;
// however, for each iterator of start time, we know we are creating a new
// rental, so usedLaptops count is always incremented
// Big O: O(n * log(n)) time | O(n) space
// only sorting happens with O(nlog(n)) time, then 1 O(n) time iteration
// should be slightly more performant than solution 1 in terms of time complexity
// func LaptopRentals(times [][2]int) int {
// 	if len(times) == 0 {
// 		return 0
// 	}

// 	usedLaptops := 0

// 	startTimes := make([]int, len(times))
// 	endtimes := make([]int, len(times))
// 	for i, interval := range times {
// 		startTimes[i] = interval[0]
// 		endtimes[i] = interval[1]
// 	}
// 	sort.Ints(startTimes)
// 	sort.Ints(endtimes)

// 	for startIterator, endIterator := 0, 0; startIterator < len(times); startIterator++ {
// 		if endtimes[endIterator] <= startTimes[startIterator] {
// 			usedLaptops -= 1
// 			endIterator += 1
// 		}
// 		usedLaptops += 1
// 	}

// 	return usedLaptops
// }
