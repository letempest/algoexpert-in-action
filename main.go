// golang entry for each AlgoExpert challenge
// #[Heaps]/Sort K-Sorted Array

package main

import (
	"fmt"

	sort "github.com/letempest/algoexpert-in-action/Heaps/Sort-K-Sorted-Array"
)

func main() {
	array := []int{3, 2, 1, 5, 4, 7, 6, 5}
	sort.SortKSortedArray(array, 3) // sorted in place
	fmt.Println(array)
}
