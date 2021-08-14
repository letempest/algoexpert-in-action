// Given a SORTED array with DISTINCT intergers, construct a BST with minimum height (or depth)
// select the number with median index as root node, and pop its left/right child by recursively calling the function
// to generate a new subtree.

import BST from '../BinarySearchTree';

// const getMedianIdx = (array: number[]): number => {
//   if (array.length % 2 === 0) {
//     return array.length / 2 - 1;
//   } else {
//     return Math.floor(array.length / 2);
//   }
// };

// // My solution: O(n) time | O(n) space
// const minHeightBST = (array: number[]): BST | null => {
//   // base case
//   if (array.length === 0) return null;

//   const curRootIdx = getMedianIdx(array);
//   const leftSubArr = array.slice(0, curRootIdx);
//   const rightSubArr = array.slice(curRootIdx + 1);

//   const currentTree = new BST(array[curRootIdx]);
//   currentTree.left = minHeightBST(leftSubArr);
//   currentTree.right = minHeightBST(rightSubArr);

//   return currentTree;
// };

// expert's solution, slightly better than mine in space complexity by passing the same src array down
// instead of creating and passing slices arrays for each recursive call
// Big O: O(n) time | O(n) space (O(d) space for accumulated frames piling on the call stack)
const minHeightBST = (array: number[]): BST => {
  return constructMinHeightBST(array, 0, array.length - 1)!;
};

const constructMinHeightBST = (
  array: number[],
  startIdx: number,
  endIdx: number
): BST | null => {
  if (endIdx < startIdx) return null;

  const curRootIdx = Math.floor((startIdx + endIdx) / 2);
  const currentNode = new BST(array[curRootIdx]);
  currentNode.left = constructMinHeightBST(array, startIdx, curRootIdx - 1);
  currentNode.right = constructMinHeightBST(array, curRootIdx + 1, endIdx);

  return currentNode;
};

const bst = minHeightBST([1, 2, 5, 7, 10, 13, 14, 15, 22]);
console.log(bst);

console.log('bst 2: ', minHeightBST([1, 5, 7, 11, 14, 16, 21, 30]));
