package main

import (
	"fmt"
)

func main() {
	srcArray := []int{3, 5, -4, 8, 11, 1, -1, 6}
	fmt.Println(twoNumberSum(srcArray, 10))
}

// // Big O: O(n) time | O(n) space for the hashtable
func twoNumberSum(array []int, target int) [][2]int {
	hashtable := map[int]bool{}
	result := [][2]int{}
	for _, num := range array {
		neededNum := target - num
		_, exists := hashtable[neededNum]
		if exists {
			result = append(result, [2]int{num, neededNum})
		}
		hashtable[num] = true
	}
	return result
}

// Solution 2
// Big O: O(n * log(n)) time | O(1) space
// func twoNumberSum(array []int, target int) [][2]int {
// 	sort.Ints(array)
// 	result := [][2]int{}
// 	for i, j := 0, len(array)-1; i < j; {
// 		num := array[i] + array[j]
// 		if num < target {
// 			i += 1
// 		} else if num > target {
// 			j -= 1
// 		} else {
// 			result = append(result, [2]int{array[i], array[j]})
// 			i += 1
// 			j -= 1
// 		}
// 	}
// 	return result
// }
