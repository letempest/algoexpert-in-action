// Given a variation of linked list, where there's a loop inside it, find out the node from which the loop starts
// a very clever conceptual solution, draw out the list and quantify the traveling distance
// Two pointers start from head, first pointer travels one node at a time; second pointer travels at 2 times speed (skip one node at a time)
// says D = (head to loopStart), P = (loopStart to overlap node), R = (overlap node to loopStart),
// and length of the linked list as T, first pointer travels distance L, hence second pointer travels 2L
// there's relations: L = D + P; 2L = 2D + 2P = T + P; hence T = D + R + P = 2D + P => (R === D)

import { LinkNode } from '../SinglyLinkList';

// Big O: O(n) time | O(1) space
const findLoop = (head: LinkNode): LinkNode => {
  let first = head.next;
  let second = head.next!.next;
  while (first !== second) {
    first = first!.next;
    second = second!.next!.next;
  }
  // at this point, the first and second pointer overlap
  // move the first pointer back to head, and both pointer has same distance from the starting node of loop
  first = head;
  while (first !== second) {
    first = first!.next;
    second = second!.next;
  }
  return first!;
};

const head = new LinkNode(0);
head.next = new LinkNode(1);
head.next.next = new LinkNode(2);
const loopStart = new LinkNode(4);
head.next.next.next = loopStart;
loopStart.next = new LinkNode(5);
loopStart.next.next = new LinkNode(6);
loopStart.next.next.next = new LinkNode(7);
loopStart.next.next.next.next = new LinkNode(8);
loopStart.next.next.next.next.next = new LinkNode(9);
loopStart.next.next.next.next.next.next = loopStart;

//
//                     5 ->  6
//                   /        \
//  0 -> 1 -> 2 ->  4          7
//                   \        /
//                    9 <--  8
//

console.log(findLoop(head));
