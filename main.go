// golang entry for each AlgoExpert challenge
// # [Arrays]/Apartment-Hunting

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Apartment-Hunting"
)

func main() {
	blocks := make([]map[string]struct{}, 5)
	var empty struct{}
	blocks[0] = map[string]struct{}{
		"Sc": empty,
	}
	blocks[1] = map[string]struct{}{
		"G": empty,
	}
	blocks[2] = map[string]struct{}{
		"G":  empty,
		"Sc": empty,
	}
	blocks[3] = map[string]struct{}{
		"Sc": empty,
	}
	blocks[4] = map[string]struct{}{
		"Sc": empty,
		"St": empty,
	}
	reqs := []string{"G", "Sc", "St"} // requirements, G - Gym, Sc - School, St - Store

	// should yield 3, which means block at index 3 has the smallest distance to all facilities in requirement array
	fmt.Println(array.ApartmentHunting(blocks, reqs))
}
