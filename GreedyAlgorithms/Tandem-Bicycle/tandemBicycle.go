package tandembicycle

import (
	"math"
	"sort"
)

// Big O: O(n * log(n)) time, for all we do is just sorting the input arrays and iterate once after that
// O(1) space, because sorting happen in place, if the input arrays are used somewhere else and not
// ok to be mutated, then copy arrays would yield O(n) space complexity
func TandemBicycle(redShirtSpeeds, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	if fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))
	} else {
		sort.Ints(blueShirtSpeeds)
	}

	var totalSpeed float64
	for i := range redShirtSpeeds {
		rider1 := float64(redShirtSpeeds[i])
		rider2 := float64(blueShirtSpeeds[i])
		totalSpeed += math.Max(rider1, rider2)
	}
	return int(totalSpeed)
}
