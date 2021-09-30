package bt

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Big O: O(n) time | O(n) space
// for a balanced binary tree, we can get the most branches (the number of leaf node),
// which is roughly 1/2 of n, so O(n) space; due to the recursive nature of the function,
// for a single branch tree, we will have at most d = n recursive calls on the callstack,
// where d is the depth of the tree, in all space complexity will be O(n) at the worst case
func BranchSums(root *BinaryTree) []int {
	sums := make([]int, 0)
	calculateBranchSums(root, &sums, 0)
	return sums
}

func calculateBranchSums(node *BinaryTree, sums *[]int, runningSum int) {
	if node == nil {
		return
	}

	runningSum += node.Value
	// Base case, leaf node
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}

	calculateBranchSums(node.Left, sums, runningSum)
	calculateBranchSums(node.Right, sums, runningSum)
}
