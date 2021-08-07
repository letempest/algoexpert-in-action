// Given a sorted! linked list, remove all duplicate node from it
// Input:  1 -> 3 -> 4 -> 4 -> 4 -> 5 -> 6 -> 6
// Output: 1 -> 3 -> 4 -> 5 -> 6
// Big O: O(n) Time | O(1) Space

import { SinglyLinkedList } from '../../src/data-structure-implementation/SinglyLinkList';

const removeDuplicatesFromLinkedList = (
  list: SinglyLinkedList
): SinglyLinkedList => {
  let currentNode = list.head;
  let nextDistinctNode = currentNode!.next;
  while (currentNode !== null) {
    nextDistinctNode = currentNode!.next;
    while (nextDistinctNode?.value === currentNode.value) {
      nextDistinctNode = nextDistinctNode.next;
    }
    currentNode.next = nextDistinctNode;
    currentNode = nextDistinctNode;
  }

  return list;
};

const list = new SinglyLinkedList();
list.push(1);
list.push(3);
list.push(4);
list.push(4);
list.push(4);
list.push(5);
list.push(6);
list.push(6);

removeDuplicatesFromLinkedList(list);
console.log(list);
