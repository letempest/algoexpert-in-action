// Given main/sub string, wrap all substrings in the main string with underscores, if multiple substring overlap each other,
// only wrap at the beginning / end of the match

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", underScorifySubstring("testthis is a testtest to see if testestest it works", "test"))
	fmt.Printf("%#v\n", underScorifySubstring("test", "test"))
}

// Big O: O(n^2 + nm) time | O(n) space, where n, m denotes the length of main, substring respectively
// for time complexity: upper bound n*(n + m), n+m comes from strings.Index finding substring for each iteration, then multiply by n, because we're traversing a length n string
// amortize analysis, depending on different main/sub string input, e.g. back to back case in which main = "testtesttesttesttest", sub = "test"
// strings.Index find the substring really quickly, there are N/M strings.Index calls, for each call it's roughly O(2M) time operation to find a match
// so O(N/M * 2M) => O(N) time complexity for this case to generate the location array
// example 2: for main = "abcabcabctest", sub = "test", two find calls (not n), and find function generally yields O(n + m), so O(n + m) for this case
// in all, time complexity for this problem can greatly vary according to different input
func underScorifySubstring(mainStr, subStr string) string {
	locations := collapse(getLocations(mainStr, subStr))
	return underscorify(mainStr, locations)
}

func getLocations(mainStr, subStr string) [][2]int {
	locations := make([][2]int, 0)
	for idx := 0; idx < len(mainStr); {
		relativeIdx := strings.Index(mainStr[idx:], subStr)
		if relativeIdx == -1 {
			break
		}
		absIdxOfMatch := idx + relativeIdx
		locations = append(locations, [2]int{absIdxOfMatch, absIdxOfMatch + len(subStr)})
		idx = absIdxOfMatch + 1
	}
	return locations
}

func collapse(locations [][2]int) [][2]int {
	if len(locations) == 0 {
		return locations
	}
	newLocations := make([][2]int, 0)
	for i, location := range locations {
		if i > 0 && location[0] <= newLocations[len(newLocations)-1][1] {
			newLocations[len(newLocations)-1][1] = location[1]
		} else {
			newLocations = append(newLocations, location)
		}
	}
	return newLocations
}

func underscorify(mainStr string, locations [][2]int) string {
	underscoredStringSlice := make([]string, 0)
	stringIdx, locationIdx := 0, 0
	for underscorePosition, inBetweenUnderscores := 0, false; stringIdx < len(mainStr) && locationIdx < len(locations); {
		if stringIdx == locations[locationIdx][underscorePosition] {
			underscoredStringSlice = append(underscoredStringSlice, "_")
			inBetweenUnderscores = !inBetweenUnderscores
			if !inBetweenUnderscores {
				locationIdx += 1
			}
			if underscorePosition == 0 {
				underscorePosition = 1
			} else {
				underscorePosition = 0
			}
		}
		underscoredStringSlice = append(underscoredStringSlice, string(mainStr[stringIdx]))
		stringIdx += 1
	}
	// main string ends interation, but inBetweenUnderscores, append a trailing "_" then
	if locationIdx < len(locations) {
		underscoredStringSlice = append(underscoredStringSlice, "_")
	} else if stringIdx < len(mainStr) { // all "_" added, yet main string not fully iterated, append the remaining characters to the end
		underscoredStringSlice = append(underscoredStringSlice, mainStr[stringIdx:])
	}

	return strings.Join(underscoredStringSlice, "")
}
