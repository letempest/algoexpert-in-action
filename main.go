// golang entry for each AlgoExpert challenge
// # Heap/Min-Heap-Construction

package main

import (
	"fmt"

	minheap "github.com/letempest/algoexpert-in-action/Heaps/Min-Heap-Construction"
)

func main() {
	heap := minheap.New([]int{30, 102, 23, 17, 18, 9, 44, 12, 31})
	// expect to be [9, 12, 23, 17, 18, 30, 44, 102, 31]
	fmt.Println(heap.Heap)
	// heap.Insert(7)
	// heap.Remove()
	heap.Peak()
}
