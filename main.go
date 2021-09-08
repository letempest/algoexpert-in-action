// golang entry for each AlgoExpert challenge
// #[Strings]/Reverse Words In String

package main

import (
	"fmt"

	string "github.com/letempest/algoexpert-in-action/Strings/Reverse-Words-In-String"
)

func main() {
	fmt.Printf("reversed result: %#v\n", string.ReverseWordsInString("AlgoExpert  is the   best!"))
	fmt.Printf("reversed result: %#v\n", string.ReverseWordsInString("   a bc  def    ghij      !hey  xyz"))
	fmt.Printf("reversed result: %#v\n", string.ReverseWordsInString("abcdefg"))
	fmt.Printf("reversed result: %#v\n", string.ReverseWordsInString("     "))
}
