// Given an array of number representing amount of coins, return the smallest
// amount of coins which is non-constructible out of the given coins,
// e.g. for [5, 7, 1, 1, 2, 3, 22] we can create 1, 2, 3, 4 ... 19 (all coins
// can only be used once)
// Solution: sort the input array, init a largestConstructibleChange with 0
// then loop through the sorted array, for each coin value, if it's greater than
// current largestConstructibleChange plus 1, we know there's no way we can
// create the value (largestConstructibleChange + 1), e.g. [1, 1, 4], for idx 2,
// we can construct 1 and 2, yet 4 > 2+1, thus 3 is non-constructible.

package array

import (
	"sort"
)

// Big O: O(nlogn) time | O(n) space (if the input array is mutatable, then O(1) space)
func NonConstructibleChange(coins []int) int {
	sortedCoins := append([]int(nil), coins...)
	sort.Ints(sortedCoins)

	largestConstructibleChange := 0
	for _, coin := range sortedCoins {
		if coin > largestConstructibleChange+1 {
			return largestConstructibleChange + 1
		}
		largestConstructibleChange += coin
	}
	return largestConstructibleChange + 1
}
