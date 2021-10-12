package sort

import (
	"math"
)

// Solution 1
// Big O: O(n * log(n)) time | O(n * log(n)) space
// func MergeSort(array []int) []int {
// 	// base case
// 	if len(array) == 1 {
// 		return array
// 	}

// 	middleIdx := int(math.Floor(float64(len(array)) / 2))
// 	leftHalf := append([]int(nil), array[:middleIdx]...)
// 	rightHalf := append([]int(nil), array[middleIdx:]...)
// 	return mergeSortedArrays(MergeSort(leftHalf), MergeSort(rightHalf))
// }

// func mergeSortedArrays(leftHalf, rightHalf []int) []int {
// 	sortedArray := make([]int, len(leftHalf)+len(rightHalf))
// 	var i, j, k int
// 	for i < len(leftHalf) && j < len(rightHalf) {
// 		if leftHalf[i] <= rightHalf[j] {
// 			sortedArray[k] = leftHalf[i]
// 			i += 1
// 		} else {
// 			sortedArray[k] = rightHalf[j]
// 			j += 1
// 		}
// 		k += 1
// 	}
// 	for i < len(leftHalf) {
// 		sortedArray[k] = leftHalf[i]
// 		i += 1
// 		k += 1
// 	}
// 	for j < len(rightHalf) {
// 		sortedArray[k] = rightHalf[j]
// 		j += 1
// 		k += 1
// 	}
// 	return sortedArray
// }

// ============================================================
// Solution 2
// Big O: O(n * log(n)) time | O(n) space
// only create an auxiliary array of length n to do merge update alternatively
// with the original input array (as a swapping reference of the two sorted sub-arrays)
func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	auxiliaryArray := append([]int(nil), array...)
	mergedSortHelper(array, 0, len(array)-1, auxiliaryArray)
	return array
}

func mergedSortHelper(mainArray []int, startIdx, endIdx int, auxiliaryArray []int) {
	// base case, merge update happens in place; if the current
	// portion is of length 1, then it's sorted already
	if startIdx == endIdx {
		return
	}

	middleIdx := int(math.Floor(float64(startIdx+endIdx) / 2))
	mergedSortHelper(auxiliaryArray, startIdx, middleIdx, mainArray)
	mergedSortHelper(auxiliaryArray, middleIdx+1, endIdx, mainArray)
	doMerge(mainArray, startIdx, middleIdx, endIdx, auxiliaryArray)
}

func doMerge(mainArray []int, startIdx, middleIdx, endIdx int, auxiliaryArray []int) {
	k := startIdx
	i := startIdx
	j := middleIdx + 1
	for i <= middleIdx && j <= endIdx {
		if auxiliaryArray[i] <= auxiliaryArray[j] {
			mainArray[k] = auxiliaryArray[i]
			i += 1
		} else {
			mainArray[k] = auxiliaryArray[j]
			j += 1
		}
		k += 1
	}
	for i <= middleIdx {
		mainArray[k] = auxiliaryArray[i]
		i += 1
		k += 1
	}
	for j <= endIdx {
		mainArray[k] = auxiliaryArray[j]
		j += 1
		k += 1
	}
}
