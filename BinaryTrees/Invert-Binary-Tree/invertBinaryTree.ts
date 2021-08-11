import { BinaryTree } from '../BinaryTree';

class InvertBinaryTree extends BinaryTree {
  left: InvertBinaryTree | null = null;
  right: InvertBinaryTree | null = null;

  // Solution 1, iteration approach, breadth-first search through every node,
  // invert left/right child of this node, and add all of its children to the queue
  // Big O: O(n) time | O(n) space (for a balanced tree, there's a time all leaf nodes are in the queue, taking roughly n/2 of space)
  invertIter(this: InvertBinaryTree) {
    const queue: (InvertBinaryTree | null)[] = [this];
    let currentNode: InvertBinaryTree;
    while (queue.length > 0) {
      currentNode = queue.shift()!;
      if (currentNode === null) continue;
      currentNode.swapLeftAndRight();
      queue.push(currentNode.left);
      queue.push(currentNode.right);
    }
  }

  private swapLeftAndRight = (): void => {
    const prevLeft = this.left;
    this.left = this.right;
    this.right = prevLeft;
  };

  // Solution 2, recursive approach - applying recursively to all child nodes as they themself are valid subtrees
  // Big O: O(n) time | O(d) space where d is the depth of the tree, a number of d calls would be put on the callstack
  // until the leaf node recursion is resolved. and d for a balanced binary tree is equivalent to O(log(n)) space
  invertRecur(this: InvertBinaryTree): void {
    // base case, leaf node
    if (!this.left && !this.right) return;
    this.swapLeftAndRight();
    this.left?.invertRecur();
    this.right?.invertRecur();
  }
}

const node4 = new InvertBinaryTree(4);
const node8 = new InvertBinaryTree(8);
const node9 = new InvertBinaryTree(9);
node4.left = node8;
node4.right = node9;
const node5 = new InvertBinaryTree(5);
const node6 = new InvertBinaryTree(6);
const node7 = new InvertBinaryTree(7);
const node2 = new InvertBinaryTree(2);
const node3 = new InvertBinaryTree(3);
node2.left = node4;
node2.right = node5;
node3.left = node6;
node3.right = node7;
const invertTree = new InvertBinaryTree(1);
invertTree.left = node2;
invertTree.right = node3;

//                   1           |            1
//                2     3        |         3    2
//              4   5  6   7     |       7  6  5  4
//            8  9               |               9  8

// invertTree.invertIter();
invertTree.invertRecur();
console.log(invertTree);
