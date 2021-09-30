package sort

// Big O:
// Best/Avg: O(n * log(n)) time | O(log(n)) space
// worst: O(n^2) time | O(n) space, where the pivoting comparison always leads to
// two skewed subarrays for subsequent recursive calls, one with n - 1 elements,
// the other with one or zero element, for this case we'll be running n recursive calls,
// and each recursive call runs in O(n) time; for the same reason, there will be
// n recursive calls on the callstack taking up space
func QuickSort(array []int) {
	quickSortHelper(&array, 0, len(array)-1)
}

func quickSortHelper(array *[]int, startIdx, endIdx int) {
	// base case
	if startIdx >= endIdx {
		return
	}

	pivotIdx := startIdx
	leftIdx := startIdx + 1
	rightIdx := endIdx
	for leftIdx <= rightIdx {
		pivotVal := (*array)[pivotIdx]
		leftVal := (*array)[leftIdx]
		rightVal := (*array)[rightIdx]

		if leftVal > pivotVal && rightVal < pivotVal {
			swap(array, leftIdx, rightIdx)
			continue
		}
		if leftVal <= pivotVal {
			leftIdx += 1
		}
		if rightVal >= pivotVal {
			rightIdx -= 1
		}
	}
	swap(array, pivotIdx, rightIdx)

	leftSubarrayIsSmaller := rightIdx-1-startIdx < endIdx-leftIdx

	// apply the quick sort algorithm on the smaller subarray first to ensure the optimal
	// space complexity, instructor's word: "tail recursion", feels vague...
	// if this step is not check, then imagine a worst case scenario wherein the function
	// is always pivoting the subarray into n-1 and the last element, there will be
	// at most n recursive calls on the callstack => O(n) space for the worst case
	if leftSubarrayIsSmaller {
		quickSortHelper(array, startIdx, rightIdx-1)
		quickSortHelper(array, leftIdx, endIdx)
	} else {
		quickSortHelper(array, leftIdx, endIdx)
		quickSortHelper(array, startIdx, rightIdx-1)
	}
}

func swap(array *[]int, i, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
