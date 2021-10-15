// golang entry for each AlgoExpert challenge
// #[Famous Algorithms]/A* Algorithm

package main

import (
	"fmt"

	famous "github.com/letempest/algoexpert-in-action/FamousAlgorithms/A-Star-Algorithm"
)

func main() {
	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}

	path := famous.AStarAlgorithm(0, 1, 4, 3, graph) // start: [0, 1] => end: [4, 3]
	// Should return [[0, 1], [0, 0], [1, 0], [2, 0], [2, 1], [3, 1], [4, 1], [4, 2], [4, 3]]
	fmt.Println(path)
}
