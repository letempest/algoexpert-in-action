// golang entry for each AlgoExpert challenge
// # [Arrays]/Apartment-Hunting

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Monotonic-Array"
)

func main() {
	fmt.Println(array.IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001})) // true
	fmt.Println(array.IsMonotonic([]int{1, 1, 1, 1, 1, 1, 2}))                            // true
	fmt.Println(array.IsMonotonic([]int{1, 2, 3, 4, 5, 4, 3, 7, 8}))                      // false
}
