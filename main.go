// golang entry for each AlgoExpert challenge
// #[Searching]/Quickselect

package main

import (
	"fmt"

	search "github.com/letempest/algoexpert-in-action/Searching/Quickselect"
)

func main() {
	array := []int{8, 5, 2, 9, 7, 6, 3}
	// should return "5", third smallest value is 5
	fmt.Println(search.Quickselect(array, 3))
}
