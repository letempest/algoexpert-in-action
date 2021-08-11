// Given a binary tree, and a node, find out the next node to be visited in in-order traversal manner
// in-order traversal, node.left -> node -> node.right; with reference to node.parent can yield optimal solution with O(d) time | O(1) space
class FSBinaryTree {
  parent: FSBinaryTree | null = null;
  left: FSBinaryTree | null = null;
  right: FSBinaryTree | null = null;
  constructor(public value: number) {}

  getInOrderTraversalOrder = (
    node: FSBinaryTree | null = this,
    order: FSBinaryTree[] = []
  ): FSBinaryTree[] => {
    if (!node) return order;
    node.getInOrderTraversalOrder(node.left, order);
    order.push(node);
    node.getInOrderTraversalOrder(node.right, order);

    return order;
  };

  // Solution 1, do an in-order traversal of the whole tree, and find out the successor in the stored order array
  // Big O: O(n) time | O(n) space
  // static findSuccessor = (
  //   tree: FSBinaryTree,
  //   node: FSBinaryTree
  // ): FSBinaryTree | void => {
  //   const inOrderTraversalOrder = tree.getInOrderTraversalOrder();
  //   const successorIdx = inOrderTraversalOrder.findIndex(n => n === node) + 1;
  //   if (successorIdx === 0 || successorIdx === inOrderTraversalOrder.length)
  //     return;
  //   return inOrderTraversalOrder[successorIdx];
  // };

  // Solution 2, observation 1: if the node has a right subtree, then the successor must be the furthest-left node in its right sub-tree
  // observation 2: if the node has no right subtree, its successor can't be in left subtree even if there is (for the in-order traversal works
  // in left -> visit current node -> right order), then the successor must be its ancestor (if there is)
  // Big O: O(d) time where d is depth of the tree | O(1) space
  static findSuccessor = (node: FSBinaryTree): FSBinaryTree | null => {
    let currentNode: FSBinaryTree;
    if (node.right) {
      currentNode = node.right;
      while (currentNode.left) {
        currentNode = currentNode.left;
      }
      return currentNode;
    } else {
      currentNode = node;
      while (currentNode.parent && currentNode.parent.right === currentNode) {
        currentNode = currentNode.parent;
      }
      return currentNode.parent;
    }
  };
}

//            1
//          /   \
//         2     3
//        / \
//       4    5
//      /
//     6

const tree = new FSBinaryTree(1);
tree.left = new FSBinaryTree(2);
tree.right = new FSBinaryTree(3);
tree.left.left = new FSBinaryTree(4);
tree.left.right = new FSBinaryTree(5);
tree.left.left.left = new FSBinaryTree(6);
tree.left.parent = tree;
tree.right.parent = tree;
tree.left.left.parent = tree.left;
tree.left.right.parent = tree.left;
tree.left.left.left.parent = tree.left.left;

console.log(tree.getInOrderTraversalOrder());
// solution 1
// console.log(FSBinaryTree.findSuccessor(tree, tree.left));

// solution 2
console.log(FSBinaryTree.findSuccessor(tree.left.right));
