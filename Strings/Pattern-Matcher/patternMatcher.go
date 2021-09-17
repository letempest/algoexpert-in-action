// Given two input string, pattern and src string to find a match, pattern in the format of "xxyxxy" or "yxyx" etc,
// where "x" and "y" could represent a substring in the src string, if the src string matches such a pattern,
// find out what the pattern "x" and "y" represents and return them in an array format.
// for example, pattern "xxyxxy" and string "gogopowerrangergogopowerranger" should yield ["go", "powerranger"].
package patternmatcher

import (
	"errors"
	"strings"
)

// Big O: O(n^2 + m) time | O(n + m) space, where n, m are the length of matchingString, pattern respectively
// time complexity: we are iterating through the matching string once (n times then), for each iteration,
// we are constructing a potential match string and compare it to the original string, constructing a new
// string of length n is a O(n) time operation, so generally O(n^2) time, the additional O(m) time comes
// from iterating through the pattern string to construct the auxiliary hash table
func PatternMatcher(pattern, str string) ([]string, error) {
	if len(pattern) > len(str) {
		return []string{}, errors.New("pattern longer than match string")
	}
	newPattern := getNewPattern(pattern)
	didSwitch := newPattern[0] != string(pattern[0])

	counts := map[string]int{
		"x": 0,
		"y": 0,
	}

	firstYPos := getCountsAndFirstYPos(newPattern, counts)

	if counts["y"] != 0 {
		for lenOfX := 1; lenOfX < len(str); lenOfX++ {
			yTotalLength := len(str) - counts["x"]*lenOfX
			if yTotalLength%counts["y"] != 0 || yTotalLength/counts["y"] < 0 {
				continue
			}
			lenOfY := yTotalLength / counts["y"]
			yIdx := firstYPos * lenOfX
			x := str[:lenOfX]
			y := str[yIdx : yIdx+lenOfY]

			potentialMatch := make([]string, 0)
			for _, char := range newPattern {
				if char == "x" {
					potentialMatch = append(potentialMatch, x)
				} else {
					potentialMatch = append(potentialMatch, y)
				}
			}
			if strings.Join(potentialMatch, "") == str {
				if !didSwitch {
					return []string{x, y}, nil
				} else {
					return []string{y, x}, nil
				}
			}
		}
	} else {
		if len(str)%counts["x"] == 0 {
			lenOfX := len(str) / counts["x"]
			x := str[:lenOfX]
			potentialMatch := make([]string, 0)
			for range newPattern {
				potentialMatch = append(potentialMatch, x)
			}
			if strings.Join(potentialMatch, "") == str {
				if !didSwitch {
					return []string{x, ""}, nil
				} else {
					return []string{"", x}, nil
				}
			}
		}
	}
	return []string{}, nil
}

func getNewPattern(pattern string) []string {
	patternLetters := strings.Split(pattern, "")
	// flip pattern "x" and "y" if the pattern is started with "y"
	// this helps simplify our algorithm's logic a lot, since if
	// we flip here, later before returning the matched pattern
	// we just need to flip them back, without the need to write
	// extra check in the main logic
	if string(pattern[0]) != "x" {
		for i, letter := range patternLetters {
			if letter == "x" {
				patternLetters[i] = "y"
			} else {
				patternLetters[i] = "x"
			}
		}
	}
	return patternLetters
}

func getCountsAndFirstYPos(pattern []string, counts map[string]int) int {
	firstYPos := -1
	for i, char := range pattern {
		counts[char] += 1
		if char == "y" && firstYPos == -1 {
			firstYPos = i
		}
	}
	return firstYPos
}
