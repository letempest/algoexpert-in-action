// golang entry for each AlgoExpert challenge
// #[Arrays]/Tournament Winner

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Tournament-Winner"
)

func main() {
	competitions := [][2]string{
		{"HTML", "C#"},
		{"C#", "python"},
		{"python", "HTML"},
	}
	results := []int{0, 0, 1}

	fmt.Println(array.TournamentWinner(competitions, results)) // "python"
}
