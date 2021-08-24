package main

import "fmt"

func main() {
	ways := numberOfWaysToMakeChange(10, []int{1, 5, 10, 25})
	fmt.Printf("%v ways to change a %v dollars banknote with denominations set %v.\n", ways, 10, []int{1, 5, 10, 25})
}

// Big O: O(nd) time, where n is amount of dollars, d is number of denominations; | O(n) space for the ways array
func numberOfWaysToMakeChange(n int, denominations []int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	for _, denomination := range denominations {
		// index ranging from 0 - n, actually means amount here
		for amount := range ways {
			if denomination <= amount {
				ways[amount] += ways[amount-denomination]
			}
		}
	}
	return ways[n]
}
