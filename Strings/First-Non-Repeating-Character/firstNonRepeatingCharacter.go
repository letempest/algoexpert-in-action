package string

// Solution 1: Brute Force approach
// Big O: O(n^2) time | O(1) space
// func FirstNonRepeatingCharacter(str string) int {
// 	for i := range str {
// 		var foundDuplicate bool
// 		for j := range str {
// 			if i != j && str[i] == str[j] {
// 				foundDuplicate = true
// 			}
// 		}
// 		if !foundDuplicate {
// 			return i
// 		}
// 	}
// 	return -1
// }

// Solution 2
// Big O: O(n) time | O(1) space
// space could be O(number of unique characters in input string), yet
// the input is limited to contain only lowercase english characters,
// so the hashtable is actually bounded by O(26) space at most, which
// is a constant, so constant space
func FirstNonRepeatingCharacter(str string) int {
	charactersCount := make(map[byte]int)
	for i := range str {
		character := str[i]
		if _, exists := charactersCount[character]; !exists {
			charactersCount[character] = 0
		}
		charactersCount[character] += 1
	}

	for i := range str {
		character := str[i]
		if charactersCount[character] == 1 {
			return i
		}
	}
	return -1
}
