// golang entry for each AlgoExpert challenge
// #[Strings]/Pattern Matcher

package main

import (
	"fmt"

	string "github.com/letempest/algoexpert-in-action/Strings/Pattern-Matcher"
)

func main() {
	pattern := "xxyxxy"
	matchString := "gogopowerrangergogopowerranger"
	matched, err := string.PatternMatcher(pattern, matchString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", matched)
}
