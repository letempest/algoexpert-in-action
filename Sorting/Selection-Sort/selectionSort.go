package sort

// Big O: O(n^2) time | O(1) space
func SelectionSort(array []int) []int {
	// currentIdx represents the start position of the unsorted sub-array
	for currentIdx := 0; currentIdx < len(array)-1; currentIdx++ {
		smallestIdx := currentIdx
		for i := currentIdx + 1; i < len(array); i++ {
			if array[i] < array[smallestIdx] {
				smallestIdx = i
			}
		}
		if smallestIdx != currentIdx {
			swap(array, currentIdx, smallestIdx)
		}
	}
	return array
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
