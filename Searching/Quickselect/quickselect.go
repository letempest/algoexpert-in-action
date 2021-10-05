// For this specific solution, return the kth smallest value in the given array basically same logic
// as quick sort, but no recursion needed, because everytime we finish a pivoting iteration, we only
// care about one subarray and apply same logic on it, this can be done in an iterative fassion.

package search

// Big O: O(n) time | O(1) space
// Space complexity: comparing to quick sort, this algorithm does not use recursion to pivot-sorting
// the subarrays, it's just a for-loop (iteratively changing startidx, endIdx), and sorting happens
// in place, so O(1) constant space;
// Time complexity: worst case - same as quick sort, might end up pivoting into a zero/one element
// subarray and a giant n-1 elements subarray, searching in this manner could yield n*n = O(n^2) time;
// Best/Avg case - imagine the pivot index we chose always perfectly divides the current array into
// two half, different to quicksort, this algorithm doesn't care about the actual sorting, for every
// pivot iteration, we can eliminate half of the values (because the "position" does not lie in that
// range), so in math it's geometric series O(N + 1/2N + 1/4N + 1/8N + ...) time, this will converge
// into O(2N), therefore O(N) time
func Quickselect(array []int, k int) int {
	position := k - 1
	return quickselectHelper(array, 0, len(array)-1, position)
}

func quickselectHelper(array []int, startIdx, endIdx, position int) int {
	for {
		if startIdx > endIdx {
			panic("Your algorithm should never arrive here!")
		}
		pivotIdx := startIdx
		leftIdx := startIdx + 1
		rightIdx := endIdx
		for leftIdx <= rightIdx {
			if array[leftIdx] > array[pivotIdx] && array[rightIdx] < array[pivotIdx] {
				swap(array, leftIdx, rightIdx)
			}
			if array[leftIdx] <= array[pivotIdx] {
				leftIdx += 1
			}
			if array[rightIdx] >= array[pivotIdx] {
				rightIdx -= 1
			}
		}
		swap(array, pivotIdx, rightIdx)

		if rightIdx == position {
			return array[rightIdx]
		} else if position < rightIdx {
			endIdx = rightIdx - 1
		} else {
			startIdx = rightIdx + 1
		}
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
