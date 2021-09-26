// golang entry for each AlgoExpert challenge
// #[Recursion]/Lowest Common Manager

package main

import (
	"fmt"

	recursion "github.com/letempest/algoexpert-in-action/Recursion/Lowest-Common-Manager"
)

//                                          A
//                      B            C       D       E       F
//                   /  |   \        |      / \             / \
//                G     H     I      J      K  L            M  N
//                  O  P  Q  R                                                (o, p, q, r are all child of h)
//                    /\      \
//                   T  U      V
//                           / | \
//                          W  X  Y
//                             |
//                             Z

func main() {
	topManager := recursion.Hierarchy{Name: "A"}
	// course video uses list here for reports array, which makes initializing the tree super nasty
	topManager.AddReports("B", "C", "D", "E", "F")
	topManager.Reports[0].AddReports("G", "H", "I")
	topManager.Reports[1].AddReports("J")
	topManager.Reports[2].AddReports("K", "L")
	topManager.Reports[4].AddReports("M", "N")
	topManager.Reports[0].Reports[1].AddReports("O", "P", "Q", "R")
	topManager.Reports[0].Reports[1].Reports[1].AddReports("T", "U")
	topManager.Reports[0].Reports[1].Reports[3].AddReports("V")
	topManager.Reports[0].Reports[1].Reports[3].Reports[0].AddReports("W", "X", "Y")
	topManager.Reports[0].Reports[1].Reports[3].Reports[0].Reports[1].AddReports("Z")

	fmt.Printf("%+v\n", recursion.LowestCommonManager(topManager, "P", "I")) // should yield "B"
	fmt.Printf("%+v\n", recursion.LowestCommonManager(topManager, "Z", "V")) // should yield "V"
	fmt.Printf("%+v\n", recursion.LowestCommonManager(topManager, "T", "J")) // should yield "A"
}
