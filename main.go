// golang entry for each AlgoExpert challenge
// #[Recursion]/Product Sum

package main

import (
	"fmt"

	recursion "github.com/letempest/algoexpert-in-action/Recursion/Product-Sum"
)

func main() {
	array := []interface{}{5, 2, []int{7, -1}, 3, []interface{}{6, []int{-13, 8}, 4}}
	fmt.Println(recursion.ProductSum(array)) // should return integer 12
}
