// golang entry for each AlgoExpert challenge
// #[Sorting]/Quick Sort

package main

import (
	"fmt"

	sort "github.com/letempest/algoexpert-in-action/Sorting/Quick-Sort"
)

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	sort.QuickSort(array) // sorting in place
	fmt.Println(array)
}
