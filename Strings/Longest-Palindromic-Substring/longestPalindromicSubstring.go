package string

// Big O: O(n^2) time | O(1) space
// iterate through the string once, for each iteration, check whether current idx is
// the center point of a potential palindrome (odd length substring) or current idx
// is the first character to the right of mirror plane of a potential palindrome (even length);
// odd/even case can at most yield O(n/2) time operation, so total O(n^2) time
func LongestPalindromicSubstring(str string) string {
	currentLongest := [2]int{0, 1} // initialize to hold first character of the string
	for i := 1; i < len(str); i++ {
		odd := getLongestPalindromeFrom(str, i-1, i+1)
		even := getLongestPalindromeFrom(str, i-1, i)
		var longest [2]int
		if getLength(odd) > getLength(even) {
			longest = odd
		} else {
			longest = even
		}
		if getLength(longest) > getLength(currentLongest) {
			currentLongest = longest
		}
	}
	return str[currentLongest[0]:currentLongest[1]]
}

func getLongestPalindromeFrom(str string, leftIdx, rightIdx int) [2]int {
	for leftIdx >= 0 && rightIdx < len(str) {
		if str[leftIdx] != str[rightIdx] {
			break
		}
		leftIdx -= 1
		rightIdx += 1
	}
	return [2]int{leftIdx + 1, rightIdx}
}

func getLength(strRange [2]int) int {
	return strRange[1] - strRange[0]
}
