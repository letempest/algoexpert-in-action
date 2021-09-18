package waterarea

import (
	"math"
)

// Big O: O(n) time | O(n) space
// most straightforward way would use 3 for-loop, first iterating through the heights array forward
// to find out the highest pillar on the left of each index, and do the same by reverse iteration
// to find out the highest pillar on the right; now for each index we know the height of left/right
// highest pillar, says h1, h2, then minHeight = min(h1, h2) is the minimum height possible to store water,
// if minHeight is greater than pillar height at current index, the difference between them is the
// actual units of water to be stored.
func WaterArea(heights []int) int {
	maxes := make([]float64, len(heights))

	var leftMax float64
	for i, height := range heights {
		maxes[i] = leftMax
		leftMax = math.Max(leftMax, float64(height))
	}
	var rightMax float64
	for i := len(heights) - 1; i >= 0; i-- {
		height := float64(heights[i])
		minHeight := math.Min(maxes[i], rightMax) // minimum height of left and right pillar
		if height < minHeight {
			maxes[i] = minHeight - height
		} else {
			maxes[i] = 0
		}
		rightMax = math.Max(rightMax, height)
	}
	var area float64
	for _, height := range maxes {
		area += height
	}
	return int(area)
}
