package main

import (
	"fmt"
	"sort"
)

func main() {
	redShirtheights := []int{5, 8, 1, 3, 4}
	blueShirtHeights := []int{6, 9, 2, 4, 5}
	fmt.Println(classPhotos(redShirtheights, blueShirtHeights))
}

// Big O: O(n * log(n)) time, because sorting the two slices takes nlog(n) time operation, iterating yields O(n) time; | O(1) time
func classPhotos(redShirtheights, blueShirtHeights []int) bool {
	sort.SliceStable(redShirtheights, func(i, j int) bool { return redShirtheights[i] > redShirtheights[j] })
	sort.SliceStable(blueShirtHeights, func(i, j int) bool { return blueShirtHeights[i] > blueShirtHeights[j] })
	isRedShirtFirstRow := redShirtheights[0] < blueShirtHeights[0]

	for i := 0; i < len(redShirtheights); i++ {
		if isRedShirtFirstRow && redShirtheights[i] >= blueShirtHeights[i] {
			return false
		}
		if !isRedShirtFirstRow && redShirtheights[i] <= blueShirtHeights[i] {
			return false
		}
	}
	return true
}
