package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"yo", "act", "flop", "tac", "cat", "oy", "olfp"}))
}

// incomprehensible solution 1: sort each word once for nlog(n) so every anagrams would match into a single string, e.g. 'tac' / 'cat' would all be sorted into 'act'
// then sort the entire array by alphabetical order and keep track of each element's position in src array, finally pop the result array
// Big O: O(w*n*log(n) + n*w*log(w)) time, where w is words count, n is the length of longest word in the list
// n*log(n) for sorting each word in step 1; w*log(w) for sorting entire array with w words in step 2
// O(wn) space
// func groupAnagrams(words []string) [][]string {
// 	if len(words) == 0 {
// 		return [][]string{}
// 	}
// 	sortedWords := make([]string, len(words))
// 	for i, word := range words {
// 		chars := strings.Split(word, "")
// 		sort.Strings(chars)
// 		sortedWords[i] = strings.Join(chars, "")
// 	}
// 	type sortMapping struct {
// 		str string
// 		idx int
// 	}
// 	sortedSlice := make([]sortMapping, len(words))
// 	for i, val := range sortedWords {
// 		sortedSlice[i] = sortMapping{val, i}
// 	}
// 	sort.Slice(sortedSlice, func(i, j int) bool {
// 		if sortedSlice[i].str == sortedSlice[j].str {
// 			return sortedSlice[i].idx < sortedSlice[j].idx
// 		}
// 		return sortedSlice[i].str < sortedSlice[j].str
// 	})

// 	var result [][]string
// 	for i, val := range sortedSlice {
// 		if i > 0 && val.str == sortedSlice[i-1].str {
// 			result[len(result)-1] = append(result[len(result)-1], words[val.idx])
// 		} else {
// 			result = append(result, []string{words[val.idx]})
// 		}
// 	}

// 	return result
// }

// Solution 2, using hashtable
// Big O: O(w*n*log(n)) time, iterating w words, for each word with length n sorting would yield n*log(n), in total w*n*log(n)
// O(wn) space
func groupAnagrams(words []string) [][]string {
	anagrams := map[string][]string{}
	var result [][]string
	for _, word := range words {
		chars := strings.Split(word, "")
		sort.Strings(chars)
		sortedWord := strings.Join(chars, "")
		_, exists := anagrams[sortedWord]
		if !exists {
			anagrams[sortedWord] = []string{word}
		} else {
			anagrams[sortedWord] = append(anagrams[sortedWord], word)
		}
	}
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}
