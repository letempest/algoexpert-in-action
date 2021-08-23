package main

import (
	"fmt"
	"sort"
)

func main() {
	srcArray := []int{3, 2, 1, 2, 6}
	fmt.Println(minimumWaitingtime(srcArray))
}

// Big O: O(n*log(n)) time for sorting the input array | O(1) space
func minimumWaitingtime(queries []int) int {
	// Sorting in place
	sort.Slice(queries, func(i, j int) bool {
		return queries[i] < queries[j]
	})

	totalWaitingTime := 0
	for i, duration := range queries {
		queriesLeft := len(queries) - (i + 1)
		totalWaitingTime += queriesLeft * duration
	}

	return totalWaitingTime
}
