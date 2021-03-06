// Given three strings, return whether the third string can be generated by interweaving the other two strings
// interweave: one letter from either src string for one step, until all letter from both strings are used

package main

import "fmt"

func main() {
	stringOne := "aaa"
	stringTwo := "aaaf"
	stringThree := "aaafaaa"
	fmt.Println(interweavingStrings(stringOne, stringTwo, stringThree))
}

// Solution 1: naive approach, no memoization, brute-force recursion with any possible combination
// Big O: O(2^(n + m)) time | O(n + m) space, at most n + m recursive calls on the callstack
// func interweavingStrings(one, two, three string) bool {
// 	if len(one)+len(two) != len(three) {
// 		return false
// 	}

// 	return areInterwoven(one, two, three, 0, 0)
// }

// func areInterwoven(one, two, three string, i, j int) bool {
// 	k := i + j
// 	if k == len(three) {
// 		return true
// 	}

// 	if i < len(one) && one[i] == three[k] {
// 		if areInterwoven(one, two, three, i+1, j) {
// 			return true
// 		}
// 	}

// 	if j < len(two) && two[j] == three[k] {
// 		return areInterwoven(one, two, three, i, j+1)
// 	}

// 	return false
// }

// Solution 2: with memoization
// Big O: O(nm) time | O(nm) space for cache matrix, n+m recursive calls on callstack < nm
func interweavingStrings(one, two, three string) bool {
	if len(one)+len(two) != len(three) {
		return false
	}

	// Cache initialized with zeros, later on if recursive function at index i, j return true, then store 1 at cache[i][j];
	// likewise, store -1 at cache[i][j] if return false
	cache := make([][]int, len(one)+1)
	for i := range cache {
		cache[i] = make([]int, len(two)+1)
	}

	return areInterwoven(one, two, three, 0, 0, cache)
}

func areInterwoven(one, two, three string, i, j int, cache [][]int) bool {
	// for every recursive call, if cache[i][j] is not equal to 0, that means we've hit this index combination before,
	// and result was already stored in cache, then return the intermediate result directly
	if cache[i][j] != 0 {
		// fmt.Printf("cache matching, i: %v, j: %v\n", i, j)
		return cache[i][j] == 1
	}

	k := i + j
	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		result := areInterwoven(one, two, three, i+1, j, cache)
		cache[i][j] = mapBoolToCache(result)
		if result {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		result := areInterwoven(one, two, three, i, j+1, cache)
		cache[i][j] = mapBoolToCache(result)
		return result
	}

	cache[i][j] = -1
	return false
}

func mapBoolToCache(result bool) int {
	if result {
		return 1
	} else {
		return -1
	}
}
