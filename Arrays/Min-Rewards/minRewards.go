package main

import (
	"fmt"
	"math"
)

func main() {
	scores := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	fmt.Println(minRewards(scores))
}

// Solution 1: naive approach
// Big O: O(n^2) time | O(n) space
// func minRewards(scores []int) int {
// 	rewards := make([]int, len(scores))
// 	for i := range rewards {
// 		rewards[i] = 1
// 	}

// 	for i := 1; i < len(scores); i++ {
// 		score := scores[i]
// 		if score > scores[i-1] {
// 			rewards[i] = rewards[i-1] + 1
// 		} else {
// 			rewards[i] = 1
// 			for j := i - 1; j >= 0; j-- {
// 				rewards[j] = int(math.Max(float64(rewards[j]), float64(rewards[j+1]+1)))
// 				if j-1 >= 0 && scores[j-1] < scores[j] {
// 					break
// 				}
// 			}
// 		}
// 	}

// 	totalRewards := 0
// 	for _, reward := range rewards {
// 		totalRewards += reward
// 	}

// 	return totalRewards
// }

// ==============================================================
// Solution 2: Peaks and Valleys, local mins and local maxs
// O(n) time, find out all the local min indexes, then traverse left and right for each of them until next peak value, the array is only traverse once
// O(n) space
// func minRewards(scores []int) int {
// 	rewards := make([]int, len(scores))
// 	for i := range rewards {
// 		rewards[i] = 1
// 	}
// 	localMinIdxs := findLocalMinIdxs(scores)

// 	for _, localMinIdx := range localMinIdxs {
// 		traverseFromLocalMinToAdjacentPeak(localMinIdx, scores, &rewards)
// 	}

// 	totalRewards := 0
// 	for _, reward := range rewards {
// 		totalRewards += reward
// 	}
// 	return totalRewards
// }

// func findLocalMinIdxs(array []int) []int {
// 	if len(array) == 1 {
// 		return []int{0}
// 	}
// 	localMinIdxs := make([]int, 0)
// 	for i, val := range array {
// 		// edge case, first element is local min
// 		if i == 0 && val < array[i+1] {
// 			localMinIdxs = append(localMinIdxs, i)
// 		}
// 		// edge case, last element is local min
// 		if i == len(array)-1 && val < array[i-1] {
// 			localMinIdxs = append(localMinIdxs, i)
// 		}
// 		if i == 0 || i == len(array)-1 {
// 			continue
// 		}
// 		if val < array[i+1] && val < array[i-1] {
// 			localMinIdxs = append(localMinIdxs, i)
// 		}
// 	}
// 	return localMinIdxs
// }

// func traverseFromLocalMinToAdjacentPeak(localMinIdx int, scores []int, rewards *[]int) {
// 	// from valley to rightside peak
// 	for i := localMinIdx + 1; i < len(scores); i++ {
// 		if scores[i] > scores[i-1] {
// 			(*rewards)[i] = (*rewards)[i-1] + 1
// 		}
// 		if i < len(scores)-1 && scores[i+1] < scores[i] {
// 			break
// 		}
// 	}
// 	// from valley to leftside peak
// 	for i := localMinIdx - 1; i >= 0; i-- {
// 		if scores[i] > scores[i+1] {
// 			(*rewards)[i] = int(math.Max(float64((*rewards)[i]), float64((*rewards)[i+1]+1)))
// 		}
// 		if i-1 >= 0 && scores[i-1] < scores[i] {
// 			break
// 		}
// 	}
// }

// ==============================================================
// Solution 3: two iterations, left -> right, right -> left, essentially doing the same thing as local mins traversing in solution 2, but way cleaner
// O(n) time | O(n) space
func minRewards(scores []int) int {
	rewards := make([]int, len(scores))
	for i := range rewards {
		rewards[i] = 1
	}

	for i := 1; i < len(scores); i++ {
		score := scores[i]
		if score > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
		}
	}

	for i := len(scores) - 2; i >= 0; i-- {
		score := scores[i]
		if score > scores[i+1] {
			rewards[i] = int(math.Max(float64(rewards[i]), float64(rewards[i+1]+1)))
		}
	}

	totalRewards := 0
	for _, reward := range rewards {
		totalRewards += reward
	}
	return totalRewards
}
