// golang entry for each AlgoExpert challenge
// #[Arrays]/Waterfall Streams

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Waterfall-Streams"
)

func main() {
	blocks := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0},
	}
	source := 3 // water flow from blocks[0][3]
	streamsDistributionAtBottom := array.WaterfallStreams(blocks, source)
	fmt.Println(streamsDistributionAtBottom)
}
