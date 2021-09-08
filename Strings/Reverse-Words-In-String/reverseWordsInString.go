package string

// Solution 1
// Big O: O(n) time | O(n) space
// func ReverseWordsInString(str string) string {
// 	words := make([]string, 0)
// 	startOfWord := 0
// 	space := []byte(" ")[0]

// 	for idx := 0; idx < len(str); idx++ {
// 		character := str[idx]
// 		if character == space {
// 			words = append(words, str[startOfWord:idx])
// 			startOfWord = idx
// 		} else if str[startOfWord] == space {
// 			words = append(words, " ")
// 			startOfWord = idx
// 		}
// 	}
// 	words = append(words, str[startOfWord:])
// 	reverseList(&words)

// 	return strings.Join(words, "")
// }

// func reverseList(words *[]string) {
// 	for i, j := 0, len(*words)-1; i < j; {
// 		(*words)[i], (*words)[j] = (*words)[j], (*words)[i]
// 		i += 1
// 		j -= 1
// 	}
// }

// =====================================================
// Solution 2
// reverse the src string and split into an array of characters
// loop through the character array with two pointers (start/end of word),
// when characters[endOfWord] != space,it means we have found a word, so all
// that's need to do is reverse the characters ranging from startOfWord to
// endOfWord, when characters[endOfWord] == space, just don't need to do anything
// (it's meaningless to swap empty space with empty space right?)
// O(n) time | O(n) space
func ReverseWordsInString(str string) string {
	characters := []byte(str) // O(n) space
	reverseListRange(&characters, 0, len(characters)-1)

	space := []byte(" ")[0]

	// though there are two for loop, but essentially we are just looping through
	// the characters array ONCE, so O(n) time
	for startOfWord := 0; startOfWord < len(characters); {
		endOfWord := startOfWord
		for endOfWord < len(characters) && characters[endOfWord] != space {
			endOfWord += 1
		}

		// word count always smaller than n, so this function call would not yield O(n) time complexity
		reverseListRange(&characters, startOfWord, endOfWord-1)
		startOfWord = endOfWord + 1
	}
	return string(characters)
}

func reverseListRange(characters *[]byte, start, end int) {
	for start < end {
		(*characters)[start], (*characters)[end] = (*characters)[end], (*characters)[start]
		start += 1
		end -= 1
	}
}
