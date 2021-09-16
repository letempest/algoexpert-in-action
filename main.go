// golang entry for each AlgoExpert challenge
// #[Famous Algorithms]/Topological Sort

package main

import (
	"fmt"

	topology "github.com/letempest/algoexpert-in-action/FamousAlgorithms/Topological-Sort"
)

func main() {
	jobs := []int{1, 2, 3, 4}
	dependencies := [][2]int{{1, 2}, {1, 3}, {3, 2}, {4, 2}, {4, 3}}
	sortedOrder := topology.TopologicalSort(jobs, dependencies)
	fmt.Println(sortedOrder)
}
