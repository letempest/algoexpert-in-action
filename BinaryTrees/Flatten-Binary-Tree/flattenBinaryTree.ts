// Given a binary tree, flatten it into a 1D linked list-ish structure, and return the leftmost node.
// in another word, do a in-order traversal, and connect each node's left/right pointer according to the traversal order.

import { BinaryTree } from '../BinaryTree';

// Solution 1, recursive in-order traversal to get all nodes in an array,
// then iterate through the array to modify every node's left/right pointers
// Big O: O(n) time | O(n) space for the auxiliary array, and O(d) for the depth amount of calls on the callstack, while (d = log(n) < n)
// const flattenBinaryTree = (root: BinaryTree): BinaryTree => {
//   const inOrderNodes = getNodesInOrder(root);
//   for (let i = 0; i < inOrderNodes.length - 1; i++) {
//     const leftNode = inOrderNodes[i];
//     const rightNode = inOrderNodes[i + 1];
//     leftNode.right = rightNode;
//     rightNode.left = leftNode;
//   }
//   return inOrderNodes[0];
// };

// const getNodesInOrder = (
//   root: BinaryTree | null,
//   inOrderNodes: BinaryTree[] = []
// ): BinaryTree[] => {
//   if (root) {
//     getNodesInOrder(root.left, inOrderNodes);
//     inOrderNodes.push(root);
//     getNodesInOrder(root.right, inOrderNodes);
//   }
//   return inOrderNodes;
// };

// =================================================================
// Solution 2
const flattenBinaryTree = (root: BinaryTree): BinaryTree => {
  const [leftMost] = flattenTree(root);
  return leftMost;
};

// My solution, for current root node, this recursive func return left/right most node in current tree
// const flattenTree = (root: BinaryTree): [BinaryTree, BinaryTree] => {
//   // Base case, current node is a leaf node
//   if (!root.left && !root.right) {
//     return [root, root];
//   }
//   const [leftMostInLeftSubTree, rightMostInLeftSubTree] = (root.left &&
//     flattenTree(root.left)) ?? [root, null];
//   root.left = rightMostInLeftSubTree;
//   if (rightMostInLeftSubTree) {
//     rightMostInLeftSubTree.right = root;
//   }

//   const [leftMostInRightSubTree, rightMostInRightSubTree] = (root.right &&
//     flattenTree(root.right)) ?? [null, root];
//   root.right = leftMostInRightSubTree;
//   if (leftMostInRightSubTree) {
//     leftMostInRightSubTree.left = root;
//   }

//   return [leftMostInLeftSubTree, rightMostInRightSubTree];
// };

// =================================================================
// Expert's solution, much cleaner :(
// O(n) time | O(d) space, where d is the depth of the tree, at most d recursive calls on the callstack
const flattenTree = (node: BinaryTree): [BinaryTree, BinaryTree] => {
  let leftMost: BinaryTree, rightMost: BinaryTree;
  let leftSubTreeLeftMost: BinaryTree, leftSubTreeRightMost: BinaryTree;
  let rightSubTreeLeftMost: BinaryTree, rightSubTreeRightMost: BinaryTree;

  if (!node.left) {
    leftMost = node;
  } else {
    [leftSubTreeLeftMost, leftSubTreeRightMost] = flattenTree(node.left);
    leftMost = leftSubTreeLeftMost;
    connectNodes(leftSubTreeRightMost, node);
  }

  if (!node.right) {
    rightMost = node;
  } else {
    [rightSubTreeLeftMost, rightSubTreeRightMost] = flattenTree(node.right);
    rightMost = rightSubTreeRightMost;
    connectNodes(node, rightSubTreeLeftMost);
  }

  return [leftMost, rightMost];
};

const connectNodes = (left: BinaryTree, right: BinaryTree): void => {
  left.right = right;
  right.left = left;
};

const tree = new BinaryTree(1);
tree.left = new BinaryTree(2);
tree.left.left = new BinaryTree(4);
tree.left.right = new BinaryTree(5);
tree.left.right.left = new BinaryTree(7);
tree.left.right.right = new BinaryTree(8);
tree.right = new BinaryTree(3);
tree.right.left = new BinaryTree(6);

//                 1
//               /   \
//             2       3
//           /   \    /
//          4     5  6
//              /  \
//             7    8

// 4 - 2 - 7 - 5 - 8 - 1 - 6 - 3
console.log(flattenBinaryTree(tree));
