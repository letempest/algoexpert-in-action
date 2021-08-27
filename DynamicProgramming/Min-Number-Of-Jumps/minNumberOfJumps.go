// Given an array in which value of each element represents the number of jump that's possible to go forward
// e.g. array[0]=3, means from 0 index you can jump forward at most 3 place, i.e. can stop at index 1, 2, 3.
// return the min amount of jumps needed to reach the last element, starting from index 0

package main

import (
	"fmt"
	"math"
)

func main() {
	srcArray := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}
	fmt.Println(minNumberOfJumps((srcArray)))
}

// Solution 1
// Big O: O(n^2) time | O(n) space
// func minNumberOfJumps(array []int) int {
// 	jumps := make([]float64, len(array))
// 	for i := 1; i < len(array); i++ {
// 		jumps[i] = math.Inf(1)
// 	}

// 	for i := 1; i < len(array); i++ {
// 		for j := 0; j < i; j++ {
// 			if array[j]+j >= i {
// 				jumps[i] = math.Min(float64(jumps[i]), float64(jumps[j]+1))
// 			}
// 		}
// 	}
// 	return int(jumps[len(jumps)-1])
// }

// Solution 2
// loop once, at each index, update the maxReach, this is the farthest reachable index at this point, e.g. at index 3, the maxReach will be 5
// variable steps represents the available steps left before jumping, when it reaches 0, it means we must jump once more
// as we need to jump, available steps will change to be the difference between current index and maxReach
// O(n) time | O(1) space
func minNumberOfJumps(array []int) int {
	if len(array) == 1 {
		return 0
	}
	maxReach, steps, jumps := array[0], array[0], 0
	for i := 1; i < len(array)-1; i++ {
		maxReach = int(math.Max(float64(array[i]+i), float64(maxReach)))
		steps -= 1
		if steps == 0 {
			jumps += 1
			steps = maxReach - i
		}
	}
	return jumps + 1
}
