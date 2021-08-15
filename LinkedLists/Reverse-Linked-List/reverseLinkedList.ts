// Given a singly linked list, reverse it in order
// need three pointers while the middle one is initially pointing at head, traverse and update the three pointers value until mid pointer stop at null
import { LinkNode, SinglyLinkedList } from '../SinglyLinkList';

// Big O: O(n) time as just traversing the src linkedlist once | O(1) space
const reverseLinkedList = (list: SinglyLinkedList): SinglyLinkedList => {
  let p1: LinkNode | null = null;
  let p2: LinkNode | null = list.head;
  let p3: LinkNode | null;
  list.tail = list.head;
  while (p2) {
    p3 = p2.next;
    p2.next = p1;
    p1 = p2;
    p2 = p3;
  }
  list.head = p1;
  return list;
};

const linkedList = new SinglyLinkedList();
linkedList.push(0).push(1).push(2).push(3).push(4).push(5);

console.log(reverseLinkedList(linkedList));
