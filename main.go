// golang entry for each AlgoExpert challenge
// #[Famous Algorithms]/Dijkstra's Algorithm

package main

import (
	"fmt"

	famous "github.com/letempest/algoexpert-in-action/FamousAlgorithms/Dijkstras-Algorithm"
)

func main() {
	edges := [][][2]int{
		{{1, 7}},
		{{2, 6}, {3, 20}, {4, 3}},
		{{3, 14}},
		{{4, 2}},
		{},
		{},
	}
	minDistances := famous.DijkstrasAlgorithm(0, edges)
	// minimum distances from starting vertex 0 to everty other vertices
	// should return [0, 7, 13, 27, 10, -1]
	fmt.Println(minDistances)
}
