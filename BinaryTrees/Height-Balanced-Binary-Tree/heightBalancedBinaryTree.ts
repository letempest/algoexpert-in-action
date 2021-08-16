// Given a binary tree, return whether it's height balanced, definition of height balanced: for every node in this tree,
// the difference in height of left and right sub-tree is smaller than or equal to 1
import { BinaryTree } from '../BinaryTree';

class TreeInfo {
  constructor(
    public readonly isBalanced: boolean,
    public readonly height: number
  ) {}
}

const getTreeInfo = (node: BinaryTree | null): TreeInfo => {
  // base case, node is null
  if (!node) {
    return new TreeInfo(true, -1);
  }
  const leftTreeInfo = getTreeInfo(node.left);
  const rightTreeInfo = getTreeInfo(node.right);

  const isBalanced =
    leftTreeInfo.isBalanced &&
    rightTreeInfo.isBalanced &&
    Math.abs(leftTreeInfo.height - rightTreeInfo.height) <= 1;
  const height = Math.max(leftTreeInfo.height, rightTreeInfo.height) + 1;

  return new TreeInfo(isBalanced, height);
};

// Big O: O(n) time as it's making two recursive calls for each node | O(h) space where h is height of the tree
// due to the recursive nature of this algorithm, there's at most h recursive function calls on the callstack awaiting return
const isHeightBalancedBinaryTree = (tree: BinaryTree): boolean => {
  return getTreeInfo(tree).isBalanced;
};

const tree = new BinaryTree(1);
tree.left = new BinaryTree(2);
tree.left.left = new BinaryTree(4);
tree.left.right = new BinaryTree(5);
tree.left.right.left = new BinaryTree(7);
tree.left.right.right = new BinaryTree(8);
// tree.left.right.right.right = new BinaryTree(9);
// tree.left.right.right.right.right = new BinaryTree(10);
tree.right = new BinaryTree(3);
tree.right.right = new BinaryTree(6);

//              1
//            /   \
//           2     3
//         /  \     \
//        4    5     6
//            / \
//           7   8
// result: true

console.log('Height Balanced? : ', isHeightBalancedBinaryTree(tree));
