// Big O: O(n) Time | O(1) Space
package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}

	result := subarraySort(input)
	fmt.Println(result)

	input2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	result2 := subarraySort(input2)
	fmt.Println(result2)
	input3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result3 := subarraySort(input3)
	fmt.Println(result3)
}

func subarraySort(arr []int) [2]int {
	minOutOfOrder := math.Inf(1)
	maxOutOfOrder := math.Inf(-1)

	for i, v := range arr {
		if isOutOfOrder(i, v, arr) {
			minOutOfOrder = math.Min(minOutOfOrder, float64(v))
			maxOutOfOrder = math.Max(maxOutOfOrder, float64(v))
		}
	}
	if math.IsInf(minOutOfOrder, 1) {
		return [2]int{-1, -1}
	}
	if math.IsInf(maxOutOfOrder, -1) {
		return [2]int{-1, -1}
	}

	subarrayLeftIdx := 0
	for int(minOutOfOrder) >= arr[subarrayLeftIdx] {
		subarrayLeftIdx += 1
	}
	subarrayRightIdx := len(arr) - 1
	for int(maxOutOfOrder) <= arr[subarrayRightIdx] {
		subarrayRightIdx -= 1
	}

	return [2]int{subarrayLeftIdx, subarrayRightIdx}
}

func isOutOfOrder(i, num int, arr []int) bool {
	if i == 0 {
		return num > arr[i+1]
	}
	if i == len(arr)-1 {
		return num < arr[i-1]
	}
	return num > arr[i+1] || num < arr[i-1]
}
