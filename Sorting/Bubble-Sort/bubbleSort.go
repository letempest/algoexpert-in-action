package sort

// My solution, same space/time complexity
// func BubbleSort(array []int) []int {
// 	end := len(array) - 1
// 	for isSorted := bubbleSortHelper(&array, end); !isSorted; {
// 		isSorted = bubbleSortHelper(&array, end-1)
// 		end -= 1
// 	}
// 	return array
// }

// func bubbleSortHelper(array *[]int, end int) bool {
// 	isSorted := true
// 	for i := 0; i < end; i++ {
// 		if (*array)[i] > (*array)[i+1] {
// 			(*array)[i], (*array)[i+1] = (*array)[i+1], (*array)[i]
// 			isSorted = false
// 		}
// 	}
// 	return isSorted
// }

// Expert solution
// Big O: O(n^2) time | O(1) space
func BubbleSort(array []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				swap(i, i+1, &array)
				isSorted = false
			}
		}
		counter += 1
	}
	return array
}

func swap(i, j int, array *[]int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
