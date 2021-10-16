// golang entry for each AlgoExpert challenge
// #[Dynamic Programming]/Disk Stacking

package main

import (
	"fmt"

	dp "github.com/letempest/algoexpert-in-action/DynamicProgramming/Disk-Stacking"
)

func main() {
	disks := [][3]int{
		{2, 3, 4},
		{2, 2, 1},
		{2, 1, 2},
		{4, 4, 5},
		{2, 2, 8},
		{3, 2, 3},
	}
	stack := dp.DiskStacking(disks)
	fmt.Println(stack)
}
