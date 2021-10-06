package runlengthencode

import (
	"fmt"
	"strings"
)

// Big O: O(n) time | O(n) space, where n is the length of the input string
func RunLengthEncoding(str string) string {
	encodedStringCharacters := make([]string, 0)
	runningLength := 1
	for i := 1; i < len(str); i++ {
		currentCharacter := str[i]
		previousCharacter := str[i-1]

		if runningLength == 9 || currentCharacter != previousCharacter {
			encodedStringCharacters = append(encodedStringCharacters, fmt.Sprint(runningLength), string(previousCharacter))
			runningLength = 0
		}

		runningLength += 1
	}
	encodedStringCharacters = append(encodedStringCharacters, fmt.Sprint(runningLength), string(str[len(str)-1]))

	return strings.Join(encodedStringCharacters, "")
}
