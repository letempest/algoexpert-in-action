package main

import "fmt"

type BST struct {
	left  *BST
	right *BST
	value int
}

func main() {
	bst := BST{value: 15}
	bst.left = &BST{value: 5}
	bst.left.left = &BST{value: 2}
	bst.left.right = &BST{value: 5}
	bst.left.left.left = &BST{value: 1}
	bst.left.left.right = &BST{value: 3}
	bst.right = &BST{value: 20}
	bst.right.left = &BST{value: 17}
	bst.right.right = &BST{value: 22}

	fmt.Printf("%#v\n", findKthLargestValueInBST(&bst, 3))
}

// Solution 1: in-order traversal to get ALL node values in sorted manner, then return the Kth one counted from the end
// Big O: O(n) time | O(n) space
// func findKthLargestValueInBST(tree *BST, k int) int {
// 	sortedNodeValues := make([]int, 0)
// 	inOrderTraverse(tree, &sortedNodeValues)
// 	return sortedNodeValues[len(sortedNodeValues)-k]
// }

// func inOrderTraverse(tree *BST, sortedNodeValues *[]int) {
// 	if tree == nil {
// 		return
// 	}

// 	inOrderTraverse(tree.left, sortedNodeValues)
// 	*sortedNodeValues = append(*sortedNodeValues, tree.value)
// 	inOrderTraverse(tree.right, sortedNodeValues)
// }

// =========================================================
// Solution 2: Reverse In-order Traversal, find the largest one first, then traversing back until the Kth node is visited
type treeInfo struct {
	numberOfNodesVisited   int
	latestVisitedNodeValue int
}

// Big O: O(h + k) time, where h is height (or depth) of the bst, we're at most going to visit K nodes, but before visiting,
// we have to go down the tree with a height h to find the rightmost node first, thus come the O(h + k) time
// O(h) space, again due to the recursive nature of this algorithm, there're at most h calls on the callstack before resolving
func findKthLargestValueInBST(tree *BST, k int) int {
	treeInfo := treeInfo{0, -1}
	reverseInOrderTraverse(tree, &treeInfo, k)
	return treeInfo.latestVisitedNodeValue
}

func reverseInOrderTraverse(tree *BST, treeInfo *treeInfo, k int) {
	if tree == nil || treeInfo.numberOfNodesVisited >= k {
		return
	}

	reverseInOrderTraverse(tree.right, treeInfo, k)
	// after visiting the right node, the visited node count has actually been updated once,
	// so check whether the count has reached k already, only visit current tree node and invoke subsequent call on left node if we haven't visited k nodes at this time
	if treeInfo.numberOfNodesVisited < k {
		treeInfo.numberOfNodesVisited += 1
		treeInfo.latestVisitedNodeValue = tree.value
		reverseInOrderTraverse(tree.left, treeInfo, k)
	}
}
