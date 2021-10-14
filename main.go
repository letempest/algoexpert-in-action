// golang entry for each AlgoExpert challenge
// #[Strings]/First Non-Repeating Character

package main

import (
	"fmt"

	string "github.com/letempest/algoexpert-in-action/Strings/First-Non-Repeating-Character"
)

func main() {
	fmt.Println(string.FirstNonRepeatingCharacter("coolcode")) // index 3
	fmt.Println(string.FirstNonRepeatingCharacter("abcdcaf"))  // index 1
}
