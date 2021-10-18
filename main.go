// golang entry for each AlgoExpert challenge
// #[Arrays]/Largest Range

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Largest-Range"
)

func main() {
	arr := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	// Should return [0, 7], number range 0 - 7 are all in the array
	fmt.Println(array.LargestRange(arr))
}
