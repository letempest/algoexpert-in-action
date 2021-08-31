// #    1  2  3  4  5  6  7  8   9   10  11  12
// val  0  1  1  2  3  5  8  13  21  34  55  89

package main

import "fmt"

func main() {
	fmt.Println(getNthFib(12)) // yields 89
}

// Solution 1: vanilla recursive approach
// Big O: O(2^n) time | O(n) space, at most n recursive calls on the callstack
// func getNthFib(n int) int {
// 	if n == 1 {
// 		return 0
// 	}
// 	if n == 2 {
// 		return 1
// 	}
// 	return getNthFib(n-1) + getNthFib(n-2)
// }

// Solution 2: memoization
// Big O: O(n) time, every fibonacci number is calculated once | O(n) space, at most n recursive calls and caching hashtable
// func getNthFib(n int) int {
// 	if n == 1 {
// 		return 0
// 	}
// 	if n == 2 {
// 		return 1
// 	}
// 	memoized := map[int]int{
// 		1: 0,
// 		2: 1,
// 	}
// 	return fibonacciHelper(n, memoized)
// }

// func fibonacciHelper(n int, memoized map[int]int) int {
// 	val, exists := memoized[n]
// 	if exists {
// 		return val
// 	} else {
// 		memoized[n] = fibonacciHelper(n-1, memoized) + fibonacciHelper(n-2, memoized)
// 		return memoized[n]
// 	}
// }

// Solution 3: Iterative Approach
// O(n) time | O(1) space
func getNthFib(n int) int {
	if n == 1 {
		return 0
	}

	lastTwo := [2]int{0, 1} // rolling update, only store the n-2, n-1 fibonacci number
	var next int
	for i := 3; i < n+1; i++ {
		next = lastTwo[0] + lastTwo[1]
		lastTwo[0], lastTwo[1] = lastTwo[1], next
	}
	return lastTwo[1]
}
