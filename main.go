// golang entry for each AlgoExpert challenge
// #[Binary Trees]/Branch Sums

package main

import (
	"fmt"

	bt "github.com/letempest/algoexpert-in-action/BinaryTrees/Branch-Sums"
)

func main() {
	tree := &bt.BinaryTree{
		Value: 1,
	}
	tree.Left = &bt.BinaryTree{Value: 2}
	tree.Left.Left = &bt.BinaryTree{Value: 4}
	tree.Left.Left.Left = &bt.BinaryTree{Value: 8}
	tree.Left.Left.Right = &bt.BinaryTree{Value: 9}
	tree.Left.Right = &bt.BinaryTree{Value: 5}
	tree.Left.Right.Left = &bt.BinaryTree{Value: 10}
	tree.Right = &bt.BinaryTree{Value: 3}
	tree.Right.Left = &bt.BinaryTree{Value: 6}
	tree.Right.Right = &bt.BinaryTree{Value: 7}

	// should print [15, 16, 18, 10, 11]
	fmt.Println(bt.BranchSums(tree))
}
