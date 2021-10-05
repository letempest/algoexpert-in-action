package dp

import (
	"math"
)

// Big O: O(nc) time | O(nc) space
// where n is the number of items; c is the capacity of the knapsack
func KnapsackProblem(items [][2]int, capacity int) (int, []int) {
	// initialization of the auxiliary matrix
	knapsackValues := make([][]int, len(items)+1)
	for i := 0; i < len(knapsackValues); i++ {
		knapsackValues[i] = make([]int, capacity+1)
	}

	for i := 1; i < len(knapsackValues); i++ {
		for c := 1; c < capacity+1; c++ {
			currentItemValue, currentItemWeight := items[i-1][0], items[i-1][1]
			if currentItemWeight <= c {
				knapsackValues[i][c] = int(math.Max(float64(knapsackValues[i-1][c]), float64(knapsackValues[i-1][c-currentItemWeight]+currentItemValue)))
			} else {
				knapsackValues[i][c] = knapsackValues[i-1][c]
			}
		}
	}
	return knapsackValues[len(knapsackValues)-1][capacity], getKnapsackItems(knapsackValues, items)
}

func getKnapsackItems(knapsackValues [][]int, items [][2]int) []int {
	itemsInKnapsack := make([]int, 0, len(items))

	// backtracking to find out what item is added to the knapsack
	for i, c := len(knapsackValues)-1, len(knapsackValues[0])-1; i > 0 && c > 0; {
		// if current value is not equal to value of previous row
		// we are adding current item to the knapsack to cause this update
		// therefore index of current item should be pushed into result array
		if knapsackValues[i][c] != knapsackValues[i-1][c] {
			itemsInKnapsack = append(itemsInKnapsack, i-1)
			c -= items[i-1][1]
		}
		i -= 1
	}
	reverse(&itemsInKnapsack)
	return itemsInKnapsack
}

func reverse(array *[]int) {
	for i, j := 0, len(*array)-1; i < j; {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
		i += 1
		j -= 1
	}
}
