// golang entry for each AlgoExpert challenge
// #[Sorting]/Merge Sort

package main

import (
	"fmt"

	sort "github.com/letempest/algoexpert-in-action/Sorting/Merge-Sort"
)

func main() {
	fmt.Println(sort.MergeSort([]int{8, 5, 2, 9, 5, 6, 3}))
	fmt.Println(sort.MergeSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
