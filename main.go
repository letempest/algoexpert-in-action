// golang entry for each AlgoExpert challenge
// #[Dynamic Programming]/Knapsack Problem

package main

import (
	"fmt"

	dp "github.com/letempest/algoexpert-in-action/DynamicProgramming/Knapsack-Problem"
)

func main() {
	items := [][2]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}} // [item's monetary value, item's weight][]
	capacity := 10
	values, itemsInKnapsack := dp.KnapsackProblem(items, capacity)
	// should print out 10, [1, 3], which means item 1 and 3 are put into the knapsack,
	// accumulating a largest monetary value of 10
	fmt.Println(values, itemsInKnapsack)
}
