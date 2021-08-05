package main

import "fmt"

func main() {
	fmt.Println(isSubSequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
	fmt.Println(isSubSequence([]int{1, 2, 3}, []int{1, 2}))
	fmt.Println(isSubSequence([]int{3, 7, 2, 1, 4}, []int{7, 3, 1}))
	fmt.Println(isSubSequence([]int{-5, 12, 3, 7, 2, 4, 0, -28, 4, 3}, []int{12, 2, 0, 4}))
}

func isSubSequence(array []int, sequence []int) bool {
	var seqIdx int

	for _, v := range array {
		if seqIdx == len(sequence) {
			break
		}
		if v == sequence[seqIdx] {
			seqIdx++
		}
	}
	return seqIdx == len(sequence)
}

func isSubSequence2(array []int, sequence []int) bool {
	var arrIdx, seqIdx int

	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return seqIdx == len(sequence)
}
