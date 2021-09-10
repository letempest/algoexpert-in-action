package sort

// treat the front part starting with one element as a SORTED array, iterate through the
// remaining array, for each element, the sub-array in front is sorted, so try to insert the
// current number into the sorted array, the way is reversely looping through the front array
// with pointer j, if the last one in the sorted array (index j-1) is greater than current value,
// just swap the two value in place, and decrement j for next iteration
// Big O: O(n^2) time, worst case for a sorted input array in descending order | O(1) space
func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; {
			swap(array, j, j-1)
			j -= 1
		}
	}
	return array
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
