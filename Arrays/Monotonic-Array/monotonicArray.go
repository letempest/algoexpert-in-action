// Given an array of integer numbers, decide whether it is monotonically increasing or decreasing
// (can have multiple consecutive numbers filled with equal value)

package monotonic

// Solution 1
// Big O: O(n) time | O(1) space
// func IsMonotonic(array []int) bool {
// 	if len(array) <= 2 {
// 		return true
// 	}

// 	direction := array[1] - array[0]
// 	for i := 2; i < len(array); i++ {
// 		if direction == 0 {
// 			direction = array[i] - array[i-1]
// 			continue
// 		}
// 		if breakDirection(direction, array[i-1], array[i]) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func breakDirection(direction, previousInt, currentInt int) bool {
// 	difference := currentInt - previousInt
// 	return direction*difference < 0
// }

// Solution 2
// Big O: O(n) time | O(1) space
func IsMonotonic(array []int) bool {
	isNonDecreasing, isNonIncreasing := true, true
	for i := 1; i < len(array); i++ {
		difference := array[i] - array[i-1]
		if difference < 0 {
			isNonDecreasing = false
		}
		if difference > 0 {
			isNonIncreasing = false
		}
	}
	// if the array is filled with same integer, the two booleans will stay true after the for loop
	return isNonDecreasing || isNonIncreasing
}
