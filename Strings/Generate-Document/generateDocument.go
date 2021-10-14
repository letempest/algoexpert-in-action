package string

// Solution 1
// Big O: O(m * (m + n)) time | O(1) space, where m is the length of document, n is the length of characters
// func GenerateDocument(characters, document string) bool {
// 	for i := range document { // O(m)
// 		character := document[i]
// 		documentFrequency := countCharacterFrequency(character, document)     // O(m)
// 		charactersFrequency := countCharacterFrequency(character, characters) // O(n)
// 		if charactersFrequency < documentFrequency {
// 			return false
// 		}
// 	}
// 	return true
// }

// func countCharacterFrequency(character byte, target string) int {
// 	var frequency int
// 	for i := range target {
// 		if target[i] == character {
// 			frequency += 1
// 		}
// 	}
// 	return frequency
// }

// Solution 2
// Big O: O(c * (n + m)) time | O(c) space, where c is the number of unique characters in document string
// func GenerateDocument(characters, document string) bool {
// 	var empty interface{}
// 	alreadyCounted := make(map[byte]interface{})

// 	for i := range document { // O(m)
// 		character := document[i]
// 		if _, exists := alreadyCounted[character]; exists {
// 			continue
// 		}

// 		documentFrequency := countCharacterFrequency(character, document)     // O(m)
// 		charactersFrequency := countCharacterFrequency(character, characters) // O(n)
// 		if charactersFrequency < documentFrequency {
// 			return false
// 		}
// 		alreadyCounted[character] = empty
// 	}
// 	return true
// }

// func countCharacterFrequency(character byte, target string) int {
// 	var frequency int
// 	for i := range target {
// 		if target[i] == character {
// 			frequency += 1
// 		}
// 	}
// 	return frequency
// }

// Solution 3
// Big O: O(n + m) time | O(c) space
// where m is the length of document, n is the length of characters
// c is the number of unique characters in characters string
func GenerateDocument(characters, document string) bool {
	characterCounts := make(map[byte]int)

	for i := range characters {
		character := characters[i]
		if _, exists := characterCounts[character]; !exists {
			characterCounts[character] = 1
		} else {
			characterCounts[character] += 1
		}
	}

	for i := range document {
		character := document[i]
		if _, exists := characterCounts[character]; !exists || characterCounts[character] == 0 {
			return false
		}
		characterCounts[character] -= 1
	}
	return true
}
