// Given the height of a staircase and maximum steps can be taken for one time (e.g. maxSteps = 3, then you can go up 1, 2, or 3 steps)
// return the total number of ways to climb the stairs (from 0 to height)

package main

import (
	"fmt"
)

func main() {
	fmt.Println(staircaseTraversal(4, 2)) // yields 5
	fmt.Println(staircaseTraversal(5, 3)) // yields 13
}

// Solution 1: vanilla recursive approach, lots of repetitive computation, worst solution
// Big O: O(k^n) time, where k = maxSteps, n = height of the staircase; O(n) space
// e.g. for n = 5, k = 3, result Recur(5,3) = Recur(4,3) + Recur(3,3) + Recur(2,2) #here Recur(2,2) because maxSteps must be <= height, so not Recur(2,3)
// func staircaseTraversal(height, maxSteps int) int {
// 	return numberOfWaysToTop(height, maxSteps)
// }

// func numberOfWaysToTop(height, maxSteps int) int {
// 	if height <= 1 {
// 		return 1
// 	}

// 	numberOfWays := 0
// 	for step := 1; step < int(math.Min(float64(maxSteps), float64(height)))+1; step++ {
// 		numberOfWays += numberOfWaysToTop(height-step, maxSteps)
// 	}
// 	return numberOfWays
// }

// =========================================================
// Solution 2: Memoization
// O(n * k) time | O(n) space, where k = maxSteps, n = height of the staircase
// func staircaseTraversal(height, maxSteps int) int {
// 	memoizedWays := make([]int, height+1)
// 	memoizedWays[0] = 1
// 	memoizedWays[1] = 1

// 	return numberOfWaysToTop(height, maxSteps, &memoizedWays)
// }

// func numberOfWaysToTop(height, maxSteps int, memoizedWays *[]int) int {
// 	if height <= 1 {
// 		return 1
// 	}
// 	if (*memoizedWays)[height] != 0 {
// 		return (*memoizedWays)[height]
// 	}

// 	numberOfWays := 0
// 	for step := 1; step < int(math.Min(float64(maxSteps), float64(height)))+1; step++ {
// 		numberOfWays += numberOfWaysToTop(height-step, maxSteps, memoizedWays)
// 	}
// 	(*memoizedWays)[height] = numberOfWays
// 	return numberOfWays
// }

// =========================================================
// Solution 3: Iterative approach
// O(n * k) time | O(n) space, where k = maxSteps, n = height of the staircase
// func staircaseTraversal(height, maxSteps int) int {
// 	if height <= 1 {
// 		return 1
// 	}
// 	waysToTop := make([]int, height+1)
// 	waysToTop[0] = 1
// 	waysToTop[1] = 1

// 	for currentHeight := 2; currentHeight < height+1; currentHeight++ {
// 		for step := 1; step < int(math.Min(float64(maxSteps), float64(currentHeight)))+1; step++ {
// 			waysToTop[currentHeight] += waysToTop[currentHeight-step]
// 		}
// 	}
// 	return waysToTop[height]
// }

// =========================================================
// Solution 4: Sliding Window Approach, best of the four solutions
// O(n) time | O(n) space
func staircaseTraversal(height, maxSteps int) int {
	currentNumberOfWays := 0
	waysToTop := []int{1}

	for currentHeight := 1; currentHeight < height+1; currentHeight++ {
		startOfWindow := currentHeight - maxSteps - 1 // actually is the starting index of PREVIOUS window, for every iteration, subtract this one
		endOfWindows := currentHeight - 1
		if startOfWindow >= 0 {
			currentNumberOfWays -= waysToTop[startOfWindow]
		}
		// move the sliding window to right, and add the new value sliding in
		currentNumberOfWays += waysToTop[endOfWindows]
		waysToTop = append(waysToTop, currentNumberOfWays)
	}

	return waysToTop[height]
}
