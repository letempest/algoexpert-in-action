// golang entry for each AlgoExpert challenge
// #[Strings]/Generate Document

package main

import (
	"fmt"

	string "github.com/letempest/algoexpert-in-action/Strings/Generate-Document"
)

func main() {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"
	fmt.Println(string.GenerateDocument(characters, document))
}
