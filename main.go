// golang entry for each AlgoExpert challenge
// #[Arrays]/Non-Constructible Change

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Non-Constructible-Change"
)

func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	// should yield 20, can create change up to 19, 20 is not constructible
	fmt.Println(array.NonConstructibleChange(coins))
}
