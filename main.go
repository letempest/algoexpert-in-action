// golang entry for each AlgoExpert challenge
// #[Stacks]/Largest Rectangle Under Skyline

package main

import (
	"fmt"

	stack "github.com/letempest/algoexpert-in-action/Stacks/Largest-Rectangle-Under-Skyline"
)

func main() {
	buildings := []int{1, 3, 3, 2, 4, 1, 5, 3, 2}
	maxArea := stack.LargestRectangleUnderSkyline(buildings)
	fmt.Println(maxArea)
}
