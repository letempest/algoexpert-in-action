// Given a bst, do a tree traversal in three orders: in-order, pre-order, post-order
// Big O: O(n) time, cause we're traversing all nodes once | O(n) space, storing the order in an array with length n = nodes
// and additionally, O(d) (where d is depth of the tree) frames are push into the callstack for the recursive nature of the solution
// and d <= n, so (n) space after all

import BST from '../BinarySearchTree';

// In-order, left -> current -> right
const inOrderTraversal = (tree: BST | null, order: number[] = []): number[] => {
  if (tree) {
    inOrderTraversal(tree.left, order);
    order.push(tree.value);
    inOrderTraversal(tree.right, order);
  }
  return order;
};

// Pre-order, current -> left -> right
const preOrderTraversal = (
  tree: BST | null,
  order: number[] = []
): number[] => {
  if (tree) {
    order.push(tree.value);
    preOrderTraversal(tree.left, order);
    preOrderTraversal(tree.right, order);
  }
  return order;
};

// Post-order, left -> right -> current
const postOrderTraversal = (
  tree: BST | null,
  order: number[] = []
): number[] => {
  if (tree) {
    postOrderTraversal(tree.left, order);
    postOrderTraversal(tree.right, order);
    order.push(tree.value);
  }
  return order;
};

const bst = new BST(10)
  .insert(5)
  .insert(15)
  .insert(2)
  .insert(5)
  .insert(22)
  .insert(1);

console.log(bst);
console.log('in-order: ', inOrderTraversal(bst));
console.log('pre-order: ', preOrderTraversal(bst));
console.log('post-order: ', postOrderTraversal(bst));
