package palindrome

// Solution 1
// O(n^2) time | O(n) space
// creating a new string is O(n) time, and there's n iterations to repeat this
// func IsPalindrome(str string) bool {
// 	var reversedStr string
// 	for i := len(str) - 1; i >= 0; i-- {
// 		reversedStr += string(str[i])
// 	}
// 	return reversedStr == str
// }

// Solution 2
// O(n) time | O(n) space
// func IsPalindrome(str string) bool {
// 	reversedChars := make([]string, 0)
// 	for i := len(str) - 1; i >= 0; i-- {
// 		reversedChars = append(reversedChars, string(str[i]))
// 	}
// 	return strings.Join(reversedChars, "") == str
// }

// Solution 3, tail recursion
// O(n) time | O(n) space
// func IsPalindrome(str string) bool {
// 	return isPalindromeHelper(str, 0)
// }

// func isPalindromeHelper(str string, i int) bool {
// 	j := len(str) - 1 - i
// 	if i >= j {
// 		return true
// 	}
// 	if str[i] != str[j] {
// 		return false
// 	}
// 	return isPalindromeHelper(str, i+1)
// }

// Solution 4
// O(n) time | O(1) space
func IsPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
