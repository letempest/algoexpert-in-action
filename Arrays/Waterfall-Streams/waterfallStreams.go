package array

// Big O: O(w^2 * h) time | O(w) space
// where w is the width of input matrix, h is the height of input matrix(length)
// each position in the matrix is iterated once, for each iteration,
// the worst case scenario is current position is a block and position above has water,
// so the water has to be split and all indices left/right have to be checked once
// therefore O (w * h * w) time
func WaterfallStreams(array [][]int, source int) []int {
	rowAbove := make([]float64, len(array[0]))
	for i := range rowAbove {
		rowAbove[i] = float64(array[0][i])
	}
	rowAbove[source] = -1

	currentRow := make([]float64, len(array[0]))
	for row := 1; row < len(array); row++ {
		for j, val := range array[row] {
			currentRow[j] = float64(val)
		}

		for idx := 0; idx < len(rowAbove); idx++ {
			valueAbove := rowAbove[idx]
			hasWaterAbove := valueAbove < 0
			hasBlock := currentRow[idx] == 1

			if !hasWaterAbove {
				continue
			}
			if !hasBlock {
				currentRow[idx] += valueAbove
				continue
			}

			// current idx is a block, split the water above to left and right
			splitWater := valueAbove / 2

			rightIdx := idx
			for rightIdx+1 < len(rowAbove) {
				rightIdx += 1
				// water above is trapped
				if rowAbove[rightIdx] == 1 {
					break
				}
				// water flow down to current position
				if currentRow[rightIdx] != 1 {
					currentRow[rightIdx] += splitWater
					break
				}
			}

			leftIdx := idx
			for leftIdx-1 >= 0 {
				leftIdx -= 1
				// water above is trapped
				if rowAbove[leftIdx] == 1 {
					break
				}
				// water flow down to current position
				if currentRow[leftIdx] != 1 {
					currentRow[leftIdx] += splitWater
					break
				}
			}
		}

		copy(rowAbove, currentRow)
	}

	finalPercentages := make([]int, len(rowAbove))
	for i := range rowAbove {
		finalPercentages[i] = int(rowAbove[i] * -100)
	}
	return finalPercentages
}
