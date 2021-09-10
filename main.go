// golang entry for each AlgoExpert challenge
// #[Sorting]/Insertion Sort

package main

import (
	"fmt"

	sort "github.com/letempest/algoexpert-in-action/Sorting/Insertion-Sort"
)

func main() {
	fmt.Println(sort.InsertionSort([]int{8, 5, 2, 9, 5, 6, 3}))
}
