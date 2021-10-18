package bt

import (
	"math"
)

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type treeInfo struct {
	diameter float64
	height   float64
}

// Big O: O(n) time | O(d) space where d is the depth of the binary tree,
// since d recursive calls on the callstack
func BinaryTreeDiameter(tree BinaryTree) int {
	return int(getTreeInfo(&tree).diameter)
}

func getTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{
			diameter: 0,
			height:   0,
		}
	}

	leftTreeData := getTreeInfo(tree.Left)
	rightTreeData := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftTreeData.height + rightTreeData.height
	maxDiameterSoFar := math.Max(leftTreeData.diameter, rightTreeData.diameter)
	currentDiameter := math.Max(maxDiameterSoFar, longestPathThroughRoot)
	currentHeight := 1 + math.Max(leftTreeData.height, rightTreeData.height)

	return treeInfo{
		diameter: currentDiameter,
		height:   currentHeight,
	}
}
