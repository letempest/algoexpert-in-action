// golang entry for each AlgoExpert challenge
// #[Binary Search Trees]/Same BSTs

package main

import (
	"fmt"

	bst "github.com/letempest/algoexpert-in-action/BinarySearchTrees/Same-BSTs"
)

func main() {
	treeOne := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	treeTwo := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}
	fmt.Printf("With these two arrays can we construct two BSTs which are actually the same? %v\n", bst.SameBSTs(treeOne, treeTwo))
}
