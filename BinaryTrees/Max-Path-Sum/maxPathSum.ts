// Given a binary tree, find the greatest sum of any possible path in the tree
// limitation of path: NO single node in the path can be connected to more than two other nodes
// (in other words, a node cannot connect to its left, right child nodes and its parent at the same time)

import { BinaryTree } from '../BinaryTree';

class RunningInfo {
  constructor(
    public readonly maxSumAsBranch: number,
    public readonly maxPathSum: number
  ) {}
}

const getMaxSum = (tree: BinaryTree | null): RunningInfo => {
  // base case
  if (!tree) {
    return new RunningInfo(0, 0);
  }

  const { maxSumAsBranch: leftMaxSumAsBranch, maxPathSum: leftMaxPathSum } =
    getMaxSum(tree.left);
  const { maxSumAsBranch: rightMaxSumAsBranch, maxPathSum: rightMaxPathSum } =
    getMaxSum(tree.right);

  const maxChildSumAsBranch = Math.max(leftMaxSumAsBranch, rightMaxSumAsBranch);
  // sum of one branch, can be further added upon the parent node sum; maxChildSumAsBranch could be negative, if so, then don't include any child branch
  const maxSumAsBranch = Math.max(maxChildSumAsBranch + tree.value, tree.value);
  // with current node as root node, get the max sum, could be a path connecting this node and both left/right sub-tree
  // in this case, the sum should be recorded for comparison but cannot be further accumulated by the parent node
  const maxSumAsRootNode = Math.max(
    maxSumAsBranch,
    // the triangle case
    leftMaxSumAsBranch + tree.value + rightMaxSumAsBranch
  );
  // the running mps for current tree
  const runningMaxPathSum = Math.max(
    leftMaxPathSum,
    rightMaxPathSum,
    maxSumAsRootNode
  );

  return new RunningInfo(maxSumAsBranch, runningMaxPathSum);
};

// Big O: O(n) time, traversing the whole tree and make one recursive call on each node
// O(log(n)) space, for a balanced tree, there's at most depth=log(n) recursive function calls on the callstack
const maxPathSum = (tree: BinaryTree): number => {
  return getMaxSum(tree).maxPathSum;
};

const tree = new BinaryTree(1);
tree.left = new BinaryTree(2);
tree.left.left = new BinaryTree(4);
tree.left.right = new BinaryTree(5);
tree.right = new BinaryTree(3);
tree.right.left = new BinaryTree(6);
tree.right.right = new BinaryTree(7);

//                1
//              /   \
//             2      3
//           /  \    /  \
//          4    5  6    7

// expected result: 5+2+1+3+7=18
console.log(maxPathSum(tree));
