// golang entry for each AlgoExpert challenge
// #[Sorting]/Heap Sort

package main

import (
	"fmt"

	sort "github.com/letempest/algoexpert-in-action/Sorting/Heap-Sort"
)

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	sort.HeapSort(array)
	fmt.Println(array) // [2, 3, 5, 5, 6, 8, 9], sorting in place
}
