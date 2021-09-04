// Utilize sliding window technique, keep track of two bounding index (both starting from 0), left/right index which represent the start and end of smallest substring
// keep sliding the right index right until a substring is found (all characters and corresponding counts is matched, order does not matter)
// then slide the left index to attempt to shrink the substring while keeping the above mentioned requirement fulfilled
// if no longer fulfilled, slide the right index right again until it reaches the end of bigString
// could have multiple substrings found, with arithmetic right - left operation it's easy to decide the narrowest bound.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestSubstringContaining("abcd$ef$axb$c$", "$$abf"))
}

// Big O: O(b + s) time | O(b + s) space, where b and s denotes the length of bigString, smallString respectively
// looping over smallString once to create targetCharCounts hashtable, and iterating bigString once for the main getSubstringBounds logic
func smallestSubstringContaining(bigString, smallString string) string {
	targetCharCounts := getCharCounts(smallString)
	substringBounds := getSubstringBounds(bigString, targetCharCounts)
	return getStringFromBounds(bigString, substringBounds)
}

func getCharCounts(str string) map[byte]int {
	charCounts := map[byte]int{}
	for i := range str {
		increaseCharCount(str[i], &charCounts)
	}
	return charCounts
}

func getSubstringBounds(str string, targetCharCounts map[byte]int) [2]int {
	substringCharCounts := map[byte]int{}
	substringBounds := [2]int{0, int(^uint(0) >> 1)} // XOR unsigned int 0 to get uint with all bits set to 1, i.e. the maximum uint, and right shift 1 to cut it half and get the maximum int64
	numUniqueChars := len(targetCharCounts)
	numUniqueCharsDone := 0
	leftIdx, rightIdx := 0, 0

	for rightIdx < len(str) {
		rightChar := str[rightIdx]
		if _, exists := targetCharCounts[rightChar]; !exists {
			rightIdx += 1
			continue
		}
		increaseCharCount(rightChar, &substringCharCounts)
		if substringCharCounts[rightChar] == targetCharCounts[rightChar] {
			numUniqueCharsDone += 1
		}
		for numUniqueCharsDone == numUniqueChars && leftIdx <= rightIdx {
			substringBounds = getCloserBounds(leftIdx, rightIdx, substringBounds[0], substringBounds[1])
			leftChar := str[leftIdx]
			if _, exists := targetCharCounts[leftChar]; !exists {
				leftIdx += 1
				continue
			}
			if substringCharCounts[leftChar] == targetCharCounts[leftChar] {
				numUniqueCharsDone -= 1
			}
			decreaseCharCount(leftChar, &substringCharCounts)
			leftIdx += 1
		}
		rightIdx += 1
	}

	return substringBounds
}

func getCloserBounds(idx1, idx2, idx3, idx4 int) [2]int {
	if idx2-idx1 < idx4-idx3 {
		return [2]int{idx1, idx2}
	} else {
		return [2]int{idx3, idx4}
	}
}

func getStringFromBounds(bigString string, bounds [2]int) string {
	start, end := bounds[0], bounds[1]
	if end == int(^uint(0)>>1) {
		return "" // no substring found, the end index is never changed
	}
	return bigString[start : end+1]
}

func increaseCharCount(char byte, charCounts *map[byte]int) {
	if _, exists := (*charCounts)[char]; exists {
		(*charCounts)[char] += 1
	} else {
		(*charCounts)[char] = 1
	}
}

func decreaseCharCount(char byte, charCounts *map[byte]int) {
	(*charCounts)[char] -= 1
	if (*charCounts)[char] < 0 {
		(*charCounts)[char] = 0
	}
}
