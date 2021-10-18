// golang entry for each AlgoExpert challenge
// #[Binary Trees]/Binary Tree Diameter

package main

import (
	"fmt"

	bt "github.com/letempest/algoexpert-in-action/BinaryTrees/Binary-Tree-Diameter"
)

func main() {
	tree := bt.BinaryTree{
		Value: 1,
	}
	tree.Left = &bt.BinaryTree{Value: 3}
	tree.Left.Left = &bt.BinaryTree{Value: 7}
	tree.Left.Left.Left = &bt.BinaryTree{Value: 8}
	tree.Left.Left.Left.Left = &bt.BinaryTree{Value: 9}
	tree.Left.Right = &bt.BinaryTree{Value: 4}
	tree.Left.Right.Right = &bt.BinaryTree{Value: 5}
	tree.Left.Right.Right.Right = &bt.BinaryTree{Value: 6}
	tree.Right = &bt.BinaryTree{Value: 2}

	//            1
	//          /   \
	//         3     2
	//       /   \
	//      7     4
	//    /        \
	//   8          5
	//  /            \
	// 9              6

	// should return 6, 9-8-7-3-4-5-6, this path has the biggest diameter 6
	fmt.Println(bt.BinaryTreeDiameter(tree))
}
