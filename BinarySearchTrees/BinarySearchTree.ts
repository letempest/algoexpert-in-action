// In this implementation, there's no node/tree hierarchy,
// every node itself is a complete sub tree of its parent bst
export default class BST {
  left: BST | null = null;
  right: BST | null = null;
  constructor(public value: number) {}

  // Average: O(Log(n)) time | O(1) space
  // Worst: O(n) time | O(1) space
  insert(this: BST, val: number): BST {
    const newNode = new BST(val);

    let currentNode = this;
    while (true) {
      if (val < currentNode.value) {
        if (!currentNode.left) {
          currentNode.left = newNode;
          break;
        } else {
          currentNode = currentNode.left;
        }
      } else {
        if (!currentNode.right) {
          currentNode.right = newNode;
          break;
        } else {
          currentNode = currentNode.right;
        }
      }
    }
    return this;
  }

  // Average: O(Log(n)) time | O(1) space
  // Worst: O(n) time | O(1) space
  contains(this: BST, val: number): boolean {
    let currentNode: BST | null = this;
    while (currentNode) {
      if (val === currentNode.value) return true;
      if (val < currentNode.value) {
        currentNode = currentNode.left ?? null;
      } else {
        currentNode = currentNode.right ?? null;
      }
    }
    return false;
  }

  // Average: O(Log(N)) time | O(1) space
  // Worst: O(N) time | O(N) space
  remove(this: BST, val: number, parentNode?: BST): BST | void {
    let currentNode: BST | null = this;
    while (currentNode) {
      if (val < currentNode.value) {
        parentNode = currentNode;
        currentNode = currentNode.left;
      } else if (val > currentNode.value) {
        parentNode = currentNode;
        currentNode = currentNode.right;
      } else {
        // found the node to be removed, and keep track of its parent node
        if (currentNode.left && currentNode.right) {
          // the node to be removed has two child nodes
          // reassign currentNode.value = smallest value of right subtree
          currentNode.value = currentNode.right.getMinValue();
          // remove the node with smallest value from the right subtree
          currentNode.right.remove(currentNode.value, currentNode);
        } else if (!parentNode) {
          // the node to be removed is the root node of the tree, thus no parentNode
          // AND!! if we make it here, currentNode=rootNode(aka deleteNode) has only one or none child node
          if (currentNode.left) {
            // replacing all instance values(left, right, value) with that of leftnode,
            // equivalent to promoting leftnode to currentNode level⬆️
            currentNode.value = currentNode.left.value;
            currentNode.right = currentNode.left.right;
            // currentNode.left must be assgined last to ensure no overwrite occur
            currentNode.left = currentNode.left.left;
          } else if (currentNode.right) {
            currentNode.value = currentNode.right.value;
            currentNode.left = currentNode.right.left;
            currentNode.right = currentNode.right.right;
          } else {
            // edge case: removing the root node(w/out parent), and this node has no child
            // which means discarding the whole BST, nothing we can do with this implementation
            return;
          }
        } else if (parentNode!.left === currentNode) {
          // the node to be removed is the left child of its parent, and this node has 0 or 1 child
          parentNode!.left = currentNode.left ?? currentNode.right;
        } else if (parentNode!.right === currentNode) {
          // the node to be removed is the right child of its parent, and this node has 0 or 1 child
          parentNode!.right = currentNode.left ?? currentNode.right;
        }
        break;
      }
    }
    return this;
  }

  // get the minimum value from the given BST
  getMinValue(this: BST): number {
    let currentNode = this;
    while (currentNode.left) {
      currentNode = currentNode.left;
    }
    return currentNode.value;
  }
}

// initialize and insert
const bst2 = new BST(10)
  .insert(5)
  .insert(15)
  .insert(2)
  .insert(5)
  .insert(13)
  .insert(22)
  .insert(1)
  .insert(14)
  .insert(12);

// console.log(bst2);

// find
console.log(bst2.contains(19));
console.log(bst2.contains(13));

// delete
console.log(bst2.remove(10));
// console.log(new BST(123).remove(123));
