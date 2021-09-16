// some observations before solution: we are sorting in place and start from 0 to end;
// which means for a k sorted array, says for [3, 2, 1, 5, 4, 7, 6, 5] k=3, for index 0,
// the final element lies between index 0 to 3, in total k+1=4 elements that's possible
// to be in; when we are looking for index 1, index 0's value is already modified in place
// with sorted value, similarly the final element lies between index 0 to 4 (of src array),
// yet one value is used for index 0 already, so again 4 possible elements to find its value.
// solution: initialize a min heap with k elements, then iterate through the src array,
// for each iteration remove the smallest value from the heap and use it to modify the src array
// in place, then insert the value of next iteration into the heap

package sortksortedarray

import (
	minheap "github.com/letempest/algoexpert-in-action/Heaps/Min-Heap-Construction"
)

// Big O: O(n * log(k)) time | O(k) space
// time complexity: initialize a min-heap with k element, initialization of min-heap is O(k) time;
// Insertion/Removal for min-heap with k element is O(log(k));
// we are iterating through the array, for each iteration, just remove the smallest value from the
// heap and modify the src array in place with this value; and insert the next iteration value into
// the heap, in all 2*O(log(k)) operation for each iteration;
// so totally n * 2*log(k) + k time (the trailing k comes from initialization), and k <=n after all,
// 2*n*log(k) + k can be condensed into n*log(k) + n => n(log(k) + 1) => O(n * (log(k))) time
// space: min heap with k elements takes O(k) space, the src array is modified in place
func SortKSortedArray(array []int, k int) []int {
	if k > len(array)-1 {
		k = len(array) - 1
	}
	minHeapWithKElements := minheap.New((array)[:k+1])

	nextIndexToInsertElement := 0
	for idx := k + 1; idx < len(array); idx++ {
		minElement := minHeapWithKElements.Remove()
		array[nextIndexToInsertElement] = minElement
		nextIndexToInsertElement += 1

		currentElement := array[idx]
		minHeapWithKElements.Insert(currentElement)
	}
	for len(minHeapWithKElements.Heap) > 0 {
		minElement := minHeapWithKElements.Remove()
		array[nextIndexToInsertElement] = minElement
		nextIndexToInsertElement += 1
	}

	return array
}
