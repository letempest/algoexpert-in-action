// Given an array in which each element is a set storing facilities within current block, e.g.
// for [{"G"}, {"Sc"}, {"St"}], there's a Gym in block 1(index 0); a school in block 2(index 1); a store in block 3(index 2)
// decide the best block to live in, so that the maximum distance to ALL closest facilities is the smallest one among all blocks
// let's say block 1 is 0 distance from gym, 1 from school, 3 from store, then for block 1 the maximum distance is 3
// return the index of such block

package apmthunt

import (
	"math"
)

// Solution 1
// Big O: O(B^2 * R) time, where B is the number of blocks, R is the number of requirements,
// nested triple for loops, 2 time for B and 1 time for R, thus come O(B^2 * R) time
// O(B) space, an auxiliary maxDistancesAtBlocks array with length B to hold all max distances for final comparison
// func ApartmentHunting(blocks []map[string]struct{}, reqs []string) int {
// 	maxDistancesAtBlocks := make([]float64, len(blocks))
// 	for i := range maxDistancesAtBlocks {
// 		maxDistancesAtBlocks[i] = math.Inf(-1)
// 	}

// 	for i := range blocks {
// 		for _, req := range reqs {
// 			closestReqDistance := math.Inf(1)
// 			for j, neighborBlock := range blocks {
// 				if _, found := neighborBlock[req]; found {
// 					closestReqDistance = math.Min(closestReqDistance, float64(distanceBetween(i, j)))
// 				}
// 			}
// 			// fmt.Printf("Min distance from current block %v with req: %v is %v\n", i, req, closestReqDistance)
// 			maxDistancesAtBlocks[i] = math.Max(maxDistancesAtBlocks[i], closestReqDistance)
// 		}
// 	}
// 	return getIdxAtMinValue(maxDistancesAtBlocks)
// }

// func distanceBetween(i, j int) int {
// 	return int(math.Abs(float64(i - j)))
// }

// func getIdxAtMinValue(array []float64) int {
// 	idxAtMinValue := 0
// 	minValue := math.Inf(1)
// 	for i := range array {
// 		if array[i] < minValue {
// 			minValue = array[i]
// 			idxAtMinValue = i
// 		}
// 	}
// 	return idxAtMinValue
// }

// Solution 2, compute all min distances for closest facilities by iterating blocks array from both directions
// minDistances per requirement can then be computed in multiple O(B * R) operations,
// way better than triple nested for loop in solution 1, O(B * R) << O(B^2 * R)
// Big O: O(B * R) time | O(B * R) space
func ApartmentHunting(blocks []map[string]struct{}, reqs []string) int {
	minDistancesFromBlocks := getMinDistances(blocks, reqs)
	maxDistancesAtBlocks := getMaxDistancesAtBlocks(minDistancesFromBlocks)
	return getIdxAtMinValue(maxDistancesAtBlocks)
}

// O(B * R) time | O(B * R) space
func getMinDistances(blocks []map[string]struct{}, reqs []string) [][]float64 {
	// initialize len(reqs) by len(blocks) 2d array, filled with infinity value
	minDistancesFromBlocks := make([][]float64, len(reqs))
	for i := range minDistancesFromBlocks {
		minDistancesFromBlocks[i] = make([]float64, len(blocks))
		for j := range minDistancesFromBlocks[i] {
			minDistancesFromBlocks[i][j] = math.Inf(1)
		}
	}

	// iterate from left to right, O(B * R) time
	for i, req := range reqs {
		for j, currentBlock := range blocks {
			if _, found := currentBlock[req]; found {
				minDistancesFromBlocks[i][j] = 0
			} else {
				if j > 0 && minDistancesFromBlocks[i][j-1] != math.Inf(1) {
					minDistancesFromBlocks[i][j] = minDistancesFromBlocks[i][j-1] + 1
				}
			}
		}
	}

	// iterate from right to left, O(B * R) time
	for i := range reqs {
		for j := len(blocks) - 1; j >= 0; j-- {
			if j < len(blocks)-1 && minDistancesFromBlocks[i][j+1] != math.Inf(1) {
				minDistancesFromBlocks[i][j] = math.Min(minDistancesFromBlocks[i][j], minDistancesFromBlocks[i][j+1]+1)
			}
		}
	}

	return minDistancesFromBlocks
}

// O(B * R) time | O(B) space
func getMaxDistancesAtBlocks(minDistancesFromBlocks [][]float64) []float64 {
	maxDistancesAtBlocks := make([]float64, len(minDistancesFromBlocks[0])) // O(B) space

	// double for loop, O(B * R) time
	for i := 0; i < len(minDistancesFromBlocks); i++ {
		for j := 0; j < len(minDistancesFromBlocks[i]); j++ {
			maxDistancesAtBlocks[j] = math.Max(maxDistancesAtBlocks[j], minDistancesFromBlocks[i][j])
		}
	}
	return maxDistancesAtBlocks
}

// O(B) time
func getIdxAtMinValue(array []float64) int {
	idxAtMinValue := 0
	minValue := math.Inf(1)
	for i := range array {
		if array[i] < minValue {
			minValue = array[i]
			idxAtMinValue = i
		}
	}
	return idxAtMinValue
}
