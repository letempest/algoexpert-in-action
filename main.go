// golang entry for each AlgoExpert challenge
// #[Greedy Algorithms]/Valid Starting City

package main

import (
	"fmt"

	greedy "github.com/letempest/algoexpert-in-action/GreedyAlgorithms/Valid-Starting-City"
)

func main() {
	distances := []int{5, 25, 15, 10, 15}
	fuel := []int{1, 2, 1, 0, 3}
	milesPerGallon := 10
	fmt.Println(greedy.ValidStartingCity(distances, fuel, milesPerGallon))
}
