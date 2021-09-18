// golang entry for each AlgoExpert challenge
// #[Greedy Algorithms]/Tandem Bicycle

package main

import (
	"fmt"

	greedy "github.com/letempest/algoexpert-in-action/GreedyAlgorithms/Tandem-Bicycle"
)

func main() {
	redShirtSpeeds := []int{5, 5, 3, 9, 2}
	blueShirtSpeeds := []int{3, 6, 7, 2, 1}
	fastest := true
	speed := greedy.TandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest)
	fmt.Println(speed)
	speed = greedy.TandemBicycle(redShirtSpeeds, blueShirtSpeeds, !fastest)
	fmt.Println(speed)
}
