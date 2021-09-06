// golang entry for each AlgoExpert challenge
// # [Graphs]/Rectangle-Mania

package main

import (
	"fmt"

	graph "github.com/letempest/algoexpert-in-action/Graphs/Rectangle-Mania"
)

func main() {
	coords := [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1}}
	fmt.Println(graph.RectangleMania(coords))
}
