package main

import (
	"fmt"
	"math"
)

func main() {
	unsortedSrcSlice := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	result := findThreeLargestNumbers(unsortedSrcSlice)
	fmt.Println(result)
}

// Big O: O(n) time, traversing the entire array | O(1) space
func findThreeLargestNumbers(slice []int) [3]int {
	negativeInf := int(math.Inf(-1))
	threeLargestNums := &[3]int{negativeInf, negativeInf, negativeInf}
	for _, v := range slice {
		if v > threeLargestNums[2] {
			shiftAndUpdate(threeLargestNums, v, 2)
		} else if v > threeLargestNums[1] {
			shiftAndUpdate(threeLargestNums, v, 1)
		} else if v > threeLargestNums[0] {
			shiftAndUpdate(threeLargestNums, v, 0)
		}
	}
	return *threeLargestNums
}

func shiftAndUpdate(array *[3]int, num, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}
}
