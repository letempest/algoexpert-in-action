// golang entry for each AlgoExpert challenge
// #[Strings]/Longest Palindromic Substring

package main

import (
	"fmt"

	string "github.com/letempest/algoexpert-in-action/Strings/Longest-Palindromic-Substring"
)

func main() {
	lps := string.LongestPalindromicSubstring("abaxyzzyxf")
	fmt.Println(lps) // should return "xyzzyx"
}
