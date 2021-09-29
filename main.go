// golang entry for each AlgoExpert challenge
// #[Arrays]/Sorted Squared Array

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Sorted-Squared-Array"
)

func main() {
	arr := []int{-9, -6, -1, 2, 4, 12}
	fmt.Println(array.SortedSquaredArray(arr))
}
