// golang entry for each AlgoExpert challenge
// #[Heaps]/Laptop Rentals

package main

import (
	"fmt"

	heap "github.com/letempest/algoexpert-in-action/Heaps/Laptop-Rentals"
)

func main() {
	times := [][2]int{{0, 2}, {1, 4}, {4, 6}, {0, 4}, {7, 8}, {9, 11}, {3, 10}}
	numOfLaptopsNeeded := heap.LaptopRentals(times)
	fmt.Println(numOfLaptopsNeeded)
}
