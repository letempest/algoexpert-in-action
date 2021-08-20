// Given a tree, determine whether it is a valid BST

import { BinaryTree } from '../../BinaryTrees/BinaryTree';

// Big O: O(n) time, every node is visited once; | O(d) space, d is depth of the tree, d recursive calls on the callstack
const validateBST = (tree: BinaryTree): boolean => {
  return validateBstHelper(
    tree,
    Number.NEGATIVE_INFINITY,
    Number.POSITIVE_INFINITY
  );
};

const validateBstHelper = (
  tree: BinaryTree | null,
  min: number,
  max: number
): boolean => {
  if (!tree) {
    return true;
  }

  if (tree.value < min || tree.value >= max) return false;

  return (
    validateBstHelper(tree.left, min, tree.value) &&
    validateBstHelper(tree.right, tree.value, max)
  );
};

const tree = new BinaryTree(10);
tree.left = new BinaryTree(5);
tree.left.left = new BinaryTree(2);
tree.left.left.left = new BinaryTree(1);
tree.left.right = new BinaryTree(5);
tree.right = new BinaryTree(15);
tree.right.left = new BinaryTree(13);
tree.right.left.right = new BinaryTree(14);
tree.right.right = new BinaryTree(22);

// Add this node, then validate func will yield false
// tree.right.right.left = new BinaryTree(22);

console.log(validateBST(tree));
