// Given two strings, find out their longest common subsequence, for example below, lcs is "xyzw"

package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonSubsequence("zxvvyzw", "xkykzpw"))
}

// Solution 1: identical matrix for intermediate lcs computation
// Big O: (nm*min(n, m)) where n, m is the length of string1, string2 respectively, as this approach maintains a (n+1) by (m+1) 2D array for intermediate result storage
// for each matrix element, we're storing a slice of byte of current lcs characters, this's a O(min(n, m)) operation for the worst case (e.g. str1 is a valid substring of str2)
// Similarly O(nm*min(n, m)) space
// func longestCommonSubsequence(str1, str2 string) string {
// 	lcs := make([][][]byte, len(str1)+1)
// 	for i := range lcs {
// 		lcs[i] = make([][]byte, len(str2)+1)
// 	}

// 	for i := 1; i < len(lcs); i++ {
// 		for j := 1; j < len(lcs[i]); j++ {
// 			if str1[i-1] != str2[j-1] {
// 				left := lcs[i][j-1]
// 				above := lcs[i-1][j]
// 				if len(left) >= len(above) {
// 					lcs[i][j] = left
// 				} else {
// 					lcs[i][j] = above
// 				}
// 			} else {
// 				diagonal := lcs[i-1][j-1]
// 				lcs[i][j] = append(diagonal, str1[i-1])
// 			}
// 		}
// 	}
// 	return string(lcs[len(lcs)-1][len(lcs[0])-1])
// }

// Solution 2
// Big O: O(nm) time | O(nm) space
type lcsBackTrackInfo struct {
	currentLcsLength int
	prevRowIdx       int
	prevColIdx       int
	currentChar      byte
}

func longestCommonSubsequence(str1, str2 string) string {
	lcs := make([][]lcsBackTrackInfo, len(str1)+1)
	for i := range lcs {
		lcs[i] = make([]lcsBackTrackInfo, len(str2)+1)
	}

	for i := 1; i < len(lcs); i++ {
		for j := 1; j < len(lcs[i]); j++ {
			if str1[i-1] != str2[j-1] {
				left := lcs[i][j-1]
				above := lcs[i-1][j]
				if left.currentLcsLength >= above.currentLcsLength {
					lcs[i][j] = lcsBackTrackInfo{
						currentLcsLength: left.currentLcsLength,
						prevRowIdx:       i,
						prevColIdx:       j - 1,
					}
				} else {
					lcs[i][j] = lcsBackTrackInfo{
						currentLcsLength: above.currentLcsLength,
						prevRowIdx:       i - 1,
						prevColIdx:       j,
					}
				}
			} else {
				lcs[i][j] = lcsBackTrackInfo{
					currentLcsLength: lcs[i-1][j-1].currentLcsLength + 1,
					prevRowIdx:       i - 1,
					prevColIdx:       j - 1,
					currentChar:      str1[i-1],
				}
			}
		}
	}
	return buildSequence(lcs)
}

// Back-tracking the lcs matrix to build the actual lcs, one time job with O(n+m) time complexity; not multiply by nm
// significantly better than solution 1, wherein for each matrix entry, there's a O(min(n, m)) time operation
func buildSequence(lcs [][]lcsBackTrackInfo) string {
	var reversedSequence []byte
	for i, j := len(lcs)-1, len(lcs[0])-1; i > 0 && j > 0; {
		currentLcs := lcs[i][j]
		if currentLcs.currentChar != 0 {
			reversedSequence = append(reversedSequence, currentLcs.currentChar)
		}
		i = currentLcs.prevRowIdx
		j = currentLcs.prevColIdx
	}
	sequence := []byte{}
	for i := len(reversedSequence) - 1; i >= 0; i-- {
		sequence = append(sequence, reversedSequence[i])
	}

	return string(sequence)
}
