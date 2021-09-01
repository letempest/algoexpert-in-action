// caveat: this algorithm requires that for each node in the tree, there's parent attribute with which we can track back to its ancestor

package main

import "fmt"

type binaryTree struct {
	value  int
	parent *binaryTree
	left   *binaryTree
	right  *binaryTree
}

func main() {
	tree := binaryTree{value: 1}
	tree.left = &binaryTree{value: 2, parent: &tree}
	tree.left.left = &binaryTree{value: 4, parent: tree.left}
	tree.left.left.right = &binaryTree{value: 9, parent: tree.left.left}
	tree.right = &binaryTree{value: 3, parent: &tree}
	tree.right.left = &binaryTree{value: 6, parent: tree.right}
	tree.right.right = &binaryTree{value: 7, parent: tree.right}

	// for the example tree below, the callback function should print out 4, 9, 2, 1, 6, 3, 7
	iterativeInOrderTraversal(&tree, func(node *binaryTree) { fmt.Println(node.value) })
}

//             1
//           /   \
//          2     3
//        /     /   \
//       4     6     7
//        \
//         9
//

// Big O: O(n) time | O(1) space
func iterativeInOrderTraversal(tree *binaryTree, callback func(node *binaryTree)) {
	var previousNode, nextNode *binaryTree
	currentNode := tree
	for currentNode != nil {
		if previousNode == nil || previousNode == currentNode.parent { // just started from the root, or should explore left sub-tree of current node
			if currentNode.left != nil {
				nextNode = currentNode.left
			} else {
				callback(currentNode)
				if currentNode.right != nil {
					nextNode = currentNode.right
				} else {
					nextNode = currentNode.parent
				}
			}
		} else if previousNode == currentNode.left {
			callback(currentNode)
			if currentNode.right != nil {
				nextNode = currentNode.right
			} else {
				nextNode = currentNode.parent
			}
		} else {
			nextNode = currentNode.parent
		}
		previousNode = currentNode
		currentNode = nextNode
	}
}
