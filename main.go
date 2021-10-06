// golang entry for each AlgoExpert challenge
// #[Strings]/Run-Length Encoding

package main

import (
	"fmt"

	strings "github.com/letempest/algoexpert-in-action/Strings/Run-Length-Encoding"
)

func main() {
	fmt.Printf("%#v\n", strings.RunLengthEncoding("AAAAAAAAAAAAABBCCCCDD")) // "9A4A2B4C2D"
	fmt.Printf("%#v\n", strings.RunLengthEncoding("Z"))                     // "1Z"
}
