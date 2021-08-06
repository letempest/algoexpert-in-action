class LinkNode {
  next: LinkNode | null = null;

  constructor(public value: number) {}
}

class SinglyLinkedList {
  head: LinkNode | null = null;
  tail: LinkNode | null = null;
  // length: number = 0;

  push = (val: number): SinglyLinkedList => {
    const newNode = new LinkNode(val);
    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      this.tail!.next = newNode;
      this.tail = newNode;
    }
    // this.length += 1
    return this;
  };
}

export { LinkNode, SinglyLinkedList };
