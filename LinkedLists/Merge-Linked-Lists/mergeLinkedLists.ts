// Given two SORTED linked list, merge them in place, and return the merged linked list
// (in course it's return the head node of the merged linked list, while input is the head nodes of two lists)
// Similar to reversing linked list, have to have three pointers to keep reference of nodes at both lists when iterating thru them

import { LinkNode, SinglyLinkedList } from '../SinglyLinkList';

// Solution 1: Iterative
// Big O: O(n + m) where n and m are length of linkedlist one and two respectively; | O(1) space, mutating the pointers binding in place
const mergeLinkedLists = (
  listOne: SinglyLinkedList,
  listTwo: SinglyLinkedList
): SinglyLinkedList => {
  // pointer iterating thru list 1
  let p1: LinkNode | null = listOne.head;
  let p1Prev: LinkNode | null = null;
  // pointer iterating thru list 2
  let p2: LinkNode | null = listTwo.head;
  // set head of merged list, merging linked list two into one and return list one
  if (listTwo.head!.value < listOne.head!.value) {
    listOne.head = listTwo.head;
  }

  while (p1 && p2) {
    if (p1.value < p2.value) {
      p1Prev = p1;
      p1 = p1.next;
    } else {
      // p2 pointing node should come before node of p1
      if (p1Prev) {
        // general case to keep reference of current p2 binding node (except edge case when p1 is head thus p1Prev has no binding node yet)
        p1Prev.next = p2;
      }
      p1Prev = p2;
      p2 = p2.next;
      p1Prev.next = p1;
    }
  }
  if (!p1) {
    // linked list one is completely iterated
    p1Prev!.next = p2;
    listOne.tail = listTwo.tail;
  }
  // clean up pointer binding of linked list two after it is merged
  listTwo.head = null;
  listTwo.tail = null;

  return listOne;
};

// // Solution 2: Recursive
// // Big O: O(n + m) where n and m are length of linkedlist one and two respectively; | O(n + m) space, n + m recursive function calls on the callstack
// const mergeLinkedListsRecur = (
//   listOne: SinglyLinkedList,
//   listTwo: SinglyLinkedList
// ): SinglyLinkedList => {
//   const tail = recursiveMerge(listOne.head, listTwo.head);
//   listOne.head =
//     listOne.head!.value < listTwo.head!.value ? listOne.head : listTwo.head;
//   listOne.tail = tail === WhichTail.listTwo ? listTwo.tail : listOne.tail;
//   listTwo.head = null;
//   listTwo.tail = null;
//   return listOne;
// };

// enum WhichTail {
//   listOne = 0,
//   listTwo = 1
// }

// const recursiveMerge = (
//   p1: LinkNode | null,
//   p2: LinkNode | null,
//   p1Prev: LinkNode | null = null
// ): WhichTail => {
//   if (!p1) {
//     p1Prev!.next = p2;
//     return WhichTail.listTwo;
//   }
//   if (!p2) {
//     return WhichTail.listOne;
//   }
//   if (p1.value < p2.value) {
//     p1Prev = p1;
//     p1 = p1.next;
//   } else {
//     if (p1Prev) {
//       p1Prev.next = p2;
//     }
//     p1Prev = p2;
//     p2 = p2.next;
//     p1Prev.next = p1;
//   }
//   return recursiveMerge(p1, p2, p1Prev);
// };

const ll1 = new SinglyLinkedList().push(2).push(6).push(7).push(8);
const ll2 = new SinglyLinkedList()
  .push(1)
  .push(3)
  .push(4)
  .push(5)
  .push(9)
  .push(10);

// linked list One: 2 -> 6 -> 7 -> 8
// linked list Two: 1 -> 3 -> 4 -> 5 -> 9 -> 10
// merged list: 1 -> 2 -> 3 -> ... -> 9 -> 10
console.log(mergeLinkedLists(ll1, ll2));

// // Solution 2
// console.log(mergeLinkedListsRecur(ll1, ll2));
