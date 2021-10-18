package array

// Big O: O(n) time | O(n) space
func LargestRange(array []int) [2]int {
	numsMap := make(map[int]bool)
	for _, num := range array {
		numsMap[num] = true
	}
	var bestRange [2]int
	var longestLength int

	for _, num := range array {
		if !numsMap[num] {
			continue
		}
		numsMap[num] = false
		currentLength := 1

		left, right := num-1, num+1

		for ifExistThenVisit(numsMap, left) {
			currentLength += 1
			left -= 1
		}

		for ifExistThenVisit(numsMap, right) {
			currentLength += 1
			right += 1
		}

		if currentLength > longestLength {
			bestRange[0] = left + 1
			bestRange[1] = right - 1
			longestLength = currentLength
		}
	}
	return bestRange
}

func ifExistThenVisit(numsMap map[int]bool, num int) bool {
	_, exists := numsMap[num]
	if exists {
		numsMap[num] = false
	}
	return exists
}
