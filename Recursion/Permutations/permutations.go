// Given an array of number, says [1,2,3], return all possible permutations

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	fmt.Println(permutations(array))
}

// Solution 1
// Big O: Upper Bound O(n^2*n!) time, view the whole callstack as a tree, there's n! leaf nodes, to reach each of these leaf node,
// there's a path of n vertex on the chain(callstack), so total n*n! calls; for each recursive call, two O(n) time operations are performed, thus come the O(n^2*n!) time
// O(n*n!) space (n! permutations, each has a length n array)
// func permutations(array []int) [][]int {
// 	permutations := &[][]int{}
// 	permutationsHelper(array, []int{}, permutations)
// 	return *permutations
// }

// func permutationsHelper(array []int, currentPermutation []int, permutations *[][]int) {
// 	if len(array) == 0 && len(currentPermutation) > 0 {
// 		*permutations = append(*permutations, currentPermutation)
// 	} else {
// 		for i := 0; i < len(array); i++ {
// 			newArray := append(append([]int{}, array[:i]...), array[i+1:]...)
// 			newPermutation := append(append([]int{}, currentPermutation...), array[i])
// 			permutationsHelper(newArray, newPermutation, permutations)
// 		}
// 	}
// }

// Solution 2
// Big O: O(n*n!) time, n*n! recursive calls, for each call, just swapping elements in place, O(1) operation
// Only for the base case of recursive func, once in that n*n! calls, to make a length n snapshot array, it's O(n) time complexity
// so in general n!*(n-1)*1 + n!*1*n => O(n*n!) time complexity
// O(n*n!) space
func permutations(array []int) [][]int {
	permutations := [][]int{}
	permutationsHelper(0, &array, &permutations)
	return permutations
}

// array element swapping in place
func permutationsHelper(i int, array *[]int, permutations *[][]int) {
	// base case: reach the end of array, take a snapshot, it's a valid permutation
	if i == len(*array)-1 {
		snapshotArr := make([]int, len(*array))
		copy(snapshotArr, *array)
		*permutations = append(*permutations, snapshotArr)
	} else {
		for j := i; j < len(*array); j++ {
			swap(array, i, j)
			permutationsHelper(i+1, array, permutations)
			swap(array, i, j)
		}
	}
}

func swap(array *[]int, i int, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
