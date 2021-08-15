class LinkNode<T> {
  prev: LinkNode<T> | null = null;
  next: LinkNode<T> | null = null;
  constructor(public value: T) {}
}

export class DoublyLinkedList<T> {
  head: LinkNode<T> | null = null;
  tail: LinkNode<T> | null = null;

  // O(1) time | O(1) space
  setHead = (node: LinkNode<T>): void => {
    if (!this.head) {
      this.head = node;
      this.tail = node;
    } else {
      this.insertBefore(this.head, node);
    }
  };

  // O(1) time | O(1) space
  setTail = (node: LinkNode<T>): void => {
    if (!this.tail) {
      this.setHead(node);
    } else {
      this.insertAfter(this.tail, node);
    }
  };

  // O(1) time | O(1) space
  insertBefore = (node: LinkNode<T>, nodeToInsert: LinkNode<T>): void => {
    if (nodeToInsert === this.head && nodeToInsert === this.tail) return;
    this.remove(nodeToInsert);
    nodeToInsert.prev = node.prev;
    nodeToInsert.next = node;
    if (node.prev) {
      node.prev.next = nodeToInsert;
    } else {
      this.head = nodeToInsert;
    }
    node.prev = nodeToInsert;
  };

  // O(1) time | O(1) space
  insertAfter = (node: LinkNode<T>, nodeToInsert: LinkNode<T>): void => {
    if (nodeToInsert === this.head && nodeToInsert === this.tail) return;
    this.remove(nodeToInsert);
    nodeToInsert.prev = node;
    nodeToInsert.next = node.next;
    if (node.next) {
      node.next.prev = nodeToInsert;
    } else {
      this.tail = nodeToInsert;
    }
    node.next = nodeToInsert;
  };

  // O(p) time where p is the position | O(1) space
  insertAtPosition = (position: number, nodeToInsert: LinkNode<T>): void => {
    if (position === 1) {
      this.setHead(nodeToInsert);
      return;
    }
    let currentNode = this.head;
    let currentPosition = 1;
    while (currentNode && currentPosition !== position) {
      currentNode = currentNode.next;
      currentPosition += 1;
    }
    if (!currentNode) {
      this.setTail(nodeToInsert);
    } else {
      this.insertBefore(currentNode, nodeToInsert);
    }
  };

  // O(n) time | O(1) space
  removeNodesWithValue = (value: T): void => {
    let currentNode = this.head;
    while (currentNode) {
      let nodeToRemove = currentNode;
      currentNode = currentNode.next;
      if (nodeToRemove.value === value) {
        this.remove(nodeToRemove);
      }
    }
  };

  // O(1) time | O(1) space
  remove = (node: LinkNode<T>): void => {
    if (node === this.head) {
      this.head = this.head.next;
    }
    if (node === this.tail) {
      this.tail = this.tail.prev;
    }
    this.removeNodeBindings(node);
  };

  // O(n) time | O(1) space
  containsNodeWithValue = (value: T): boolean => {
    let currentNode = this.head;
    while (currentNode) {
      if (currentNode.value === value) return true;
      currentNode = currentNode.next;
    }
    return false;
  };

  private removeNodeBindings = (node: LinkNode<T>): void => {
    if (node.prev) {
      node.prev.next = node.next;
    }
    if (node.next) {
      node.next.prev = node.prev;
    }
    node.prev = null;
    node.next = null;
  };
}
