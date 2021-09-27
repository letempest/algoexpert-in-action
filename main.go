// golang entry for each AlgoExpert challenge
// #[Sorting]/Selection Sort

package main

import (
	"fmt"

	sort "github.com/letempest/algoexpert-in-action/Sorting/Selection-Sort"
)

func main() {
	srcArray := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(sort.SelectionSort(srcArray)) // should print [2, 3, 5, 5, 6, 8, 9]
}
