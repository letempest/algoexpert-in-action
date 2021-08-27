// Given a binary tree, target value and distance k, find out all nodes that is distance k away from the node with the target value in the tree
// Solution 1 is much easier to comprehend, though it has more data to keep track of; solution 2 is a recursive approach starting from the root node,
// for each current node, says leftDistance = l, which represents the distance from target node in the left sub-tree to currentnode,
// then all we need to do is find out nodes with distance k - l within the right sub-tree

import { BinaryTree } from '../BinaryTree';

// // Big O: O(n) time, breadth-first search yields O(v+e), and edges must be fewer than vertex, so O(n) in fact | O(n) space
// const findNodesDistanceK = (
//   tree: BinaryTree,
//   target: number,
//   k: number
// ): number[] => {
//   const nodesToParents = populateNodesToParents(tree, {}, null);
//   const targetNode = getNodeFromValue(target, tree, nodesToParents);
//   if (!targetNode) {
//     return [];
//   }

//   return breadthFirstSearchForNodesDistanceK(targetNode, nodesToParents, k);
// };

// const breadthFirstSearchForNodesDistanceK = (
//   targetNode: BinaryTree,
//   nodesToParents: NodesToParents,
//   k: number
// ): number[] => {
//   const queue: [BinaryTree, number][] = [[targetNode, 0]];
//   const seen = new Set<number>().add(targetNode.value);

//   while (queue.length > 0) {
//     const [currentNode, distanceFromTarget] = queue.shift()!;
//     if (distanceFromTarget === k) {
//       const nodesDistanceK = queue.map(([node]) => node.value);
//       nodesDistanceK.push(currentNode.value);
//       return nodesDistanceK;
//     }

//     const connectedNodes = [
//       currentNode.left,
//       currentNode.right,
//       nodesToParents[currentNode.value]
//     ];
//     for (const node of connectedNodes) {
//       if (node && !seen.has(node.value)) {
//         seen.add(node.value);
//         queue.push([node, distanceFromTarget + 1]);
//       }
//     }
//   }
//   return [];
// };

// interface NodesToParents {
//   [key: number]: BinaryTree | null;
// }

// const getNodeFromValue = (
//   value: number,
//   tree: BinaryTree,
//   nodesToParents: NodesToParents
// ): BinaryTree | null => {
//   if (tree.value === value) {
//     return tree;
//   }

//   const parent = nodesToParents[value];
//   if (parent) {
//     return parent.left?.value === value ? parent.left : parent.right;
//   } else {
//     return null;
//   }
// };

// const populateNodesToParents = (
//   node: BinaryTree | null,
//   nodesToParents: NodesToParents,
//   parent: BinaryTree | null
// ): NodesToParents => {
//   if (node) {
//     nodesToParents[node.value] = parent;
//     populateNodesToParents(node.left, nodesToParents, node);
//     populateNodesToParents(node.right, nodesToParents, node);
//   }
//   return nodesToParents;
// };

// =========================================================================
// Solution 2
// O(n) time | O(n) space
const findNodesDistanceK = (
  tree: BinaryTree,
  target: number,
  k: number
): number[] => {
  const nodesDistanceK: number[] = [];
  findDistanceFromNodeToTarget(tree, target, k, nodesDistanceK);
  return nodesDistanceK;
};

const findDistanceFromNodeToTarget = (
  node: BinaryTree | null,
  target: number,
  k: number,
  nodesDistanceK: number[] = []
): number => {
  if (!node) {
    return -1;
  }

  if (node.value === target) {
    // current node is the target node, look into current subtree to find any node that's distance (k - 0)
    addSubtreeNodesAtDistanceK(node, 0, k, nodesDistanceK);
    return 1;
  }

  const leftDistance = findDistanceFromNodeToTarget(
    node.left,
    target,
    k,
    nodesDistanceK
  );
  const rightDistance = findDistanceFromNodeToTarget(
    node.right,
    target,
    k,
    nodesDistanceK
  );

  if (leftDistance === k || rightDistance === k) {
    nodesDistanceK.push(node.value);
  }

  // target node in the left subtree, then look into the right subtree for any nodes with distance k-leftDistance
  if (leftDistance !== -1) {
    // leftDistance is the distance from targetNode in the left subtree to current node,
    // so distance from this targetNode to right childnode of current node is leftDistance + 1
    addSubtreeNodesAtDistanceK(node.right, leftDistance + 1, k, nodesDistanceK);
    return leftDistance + 1;
  }
  if (rightDistance !== -1) {
    addSubtreeNodesAtDistanceK(node.left, rightDistance + 1, k, nodesDistanceK);
    return rightDistance + 1;
  }

  // if we reach here, target node does not exist in this tree
  return -1;
};

const addSubtreeNodesAtDistanceK = (
  node: BinaryTree | null,
  distance: number,
  k: number,
  nodesDistanceK: number[]
): void => {
  if (!node) return;

  if (distance === k) {
    nodesDistanceK.push(node.value);
  } else {
    addSubtreeNodesAtDistanceK(node.left, distance + 1, k, nodesDistanceK);
    addSubtreeNodesAtDistanceK(node.right, distance + 1, k, nodesDistanceK);
  }
};

const tree = new BinaryTree(1);
tree.left = new BinaryTree(2);
tree.left.left = new BinaryTree(4);
tree.left.right = new BinaryTree(5);
tree.right = new BinaryTree(3);
tree.right.right = new BinaryTree(6);
tree.right.right.left = new BinaryTree(7);
tree.right.right.right = new BinaryTree(8);

//
//              1
//            /   \
//          2       3
//        /   \      \
//       4     5      6
//                  /   \
//                 7     8
//

// return all values of nodes which is distance 2 away from node with value 3, answer should be [2, 7, 8], order does not matter
console.log(findNodesDistanceK(tree, 3, 2));
