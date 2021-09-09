package stack

import (
	"math"
)

// Solution 1
// Big O: O(n^2) time | O(1) space
// func LargestRectangleUnderSkyline(buildings []int) int {
// 	maxArea := float64(0)
// 	for pillarIdx, currentHeight := range buildings {
// 		furthestLeft := pillarIdx
// 		for furthestLeft > 0 && currentHeight <= buildings[furthestLeft-1] {
// 			furthestLeft -= 1
// 		}

// 		furthestRight := pillarIdx
// 		for furthestRight < len(buildings)-1 && currentHeight <= buildings[furthestRight+1] {
// 			furthestRight += 1
// 		}

// 		areaWithCurrentBuilding := (furthestRight - furthestLeft + 1) * currentHeight
// 		maxArea = math.Max(maxArea, float64(areaWithCurrentBuilding))
// 	}
// 	return int(maxArea)
// }

// ================================================================
// Solution 2, I find it really tricky to understand
// initialze an empty stack to be filled with index of building, iterate through the
// buildings array once (actually one more index at the end), if stack is not empty
// and the top building in the stack is higher than or equal to building at current index,
// then we know we cannot expand right anymore, do the following:
// 1. set the current index as the right bound; 2. pop off the top building from stack,
// use this building's height as currentHeight to calculate area; 3. look into the stack,
// if stack is not empty, the building on top is to be the furthest left bound (with the way
// we compare and push building onto stack, element in the stack must be in strictly ascending order)
// if the stack is empty, then we can use the furthest left building in the src array as left bound;
// if the condition check on line 32 failed, the top building in stack is strictly shorter than current one,
// we know we have not reach the furthest right bound yet and can continue expanding right,
// simply push current index onto the stack and start next iteration
// O(n) time | O(n) space
func LargestRectangleUnderSkyline(buildings []int) int {
	pillarIndices := make([]int, 0) // O(n) space
	maxArea := float64(0)

	for i := 0; i < len(buildings)+1; i++ {
		leftBound, rightBound := 0, 0
		for len(pillarIndices) > 0 && (i == len(buildings) || buildings[pillarIndices[len(pillarIndices)-1]] >= buildings[i]) {
			currentHeight := buildings[pillarIndices[len(pillarIndices)-1]]
			pillarIndices = pillarIndices[:len(pillarIndices)-1]
			rightBound = i
			if len(pillarIndices) > 0 {
				leftBound = pillarIndices[len(pillarIndices)-1]
			} else {
				leftBound = -1
			}
			maxArea = math.Max(maxArea, float64((rightBound-leftBound-1)*currentHeight))
		}
		pillarIndices = append(pillarIndices, i)
	}
	return int(maxArea)
}
