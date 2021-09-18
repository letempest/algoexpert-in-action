// golang entry for each AlgoExpert challenge
// #[Dynamic Programming]/Water Area

package main

import (
	"fmt"

	dp "github.com/letempest/algoexpert-in-action/DynamicProgramming/Water-Area"
)

func main() {
	heights := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}
	area := dp.WaterArea(heights)
	fmt.Println(area)
}
