import { LinkNode, SinglyLinkedList } from '../SinglyLinkList';

// Big O: O(n) time | O(1) space
const rearrangeLinkedList = (
  list: SinglyLinkedList,
  key: number
): LinkNode | null => {
  let smallerListHead: LinkNode | null = null,
    smallerListTail: LinkNode | null = null;
  let equalListHead: LinkNode | null = null,
    equalListTail: LinkNode | null = null;
  let greaterListHead: LinkNode | null = null,
    greaterListTail: LinkNode | null = null;

  let currentNode = list.head;
  let prevNode: LinkNode | null;
  while (currentNode) {
    if (currentNode.value < key) {
      const [newHead, newTail] = growLinkedList(
        smallerListHead,
        smallerListTail,
        currentNode
      );
      smallerListHead = newHead;
      smallerListTail = newTail;
    } else if (currentNode.value > key) {
      const [newHead, newTail] = growLinkedList(
        greaterListHead,
        greaterListTail,
        currentNode
      );
      greaterListHead = newHead;
      greaterListTail = newTail;
    } else {
      const [newHead, newTail] = growLinkedList(
        equalListHead,
        equalListTail,
        currentNode
      );
      equalListHead = newHead;
      equalListTail = newTail;
    }

    prevNode = currentNode;
    currentNode = currentNode.next;
    prevNode.next = null;
  }

  const [firstHead, firstTail] = connectLinkedList(
    smallerListHead,
    smallerListTail,
    equalListHead,
    equalListTail
  );
  const [finalHead, finalTail] = connectLinkedList(
    firstHead,
    firstTail,
    greaterListHead,
    greaterListTail
  );
  list.head = finalHead;
  list.tail = finalTail;

  return finalHead;
};

const growLinkedList = (
  head: LinkNode | null,
  tail: LinkNode | null,
  currentNode: LinkNode
): [LinkNode | null, LinkNode | null] => {
  let newHead = head;
  let newTail = currentNode;

  if (!head) {
    newHead = currentNode;
  }
  if (tail) {
    tail.next = currentNode;
  }
  return [newHead, newTail];
};

const connectLinkedList = (
  listOneHead: LinkNode | null,
  listOneTail: LinkNode | null,
  listTwoHead: LinkNode | null,
  listTwoTail: LinkNode | null
): [LinkNode | null, LinkNode | null] => {
  const head = !listOneHead ? listTwoHead : listOneHead;
  const tail = !listTwoTail ? listOneTail : listTwoTail;

  if (listOneTail) {
    listOneTail.next = listTwoHead;
  }
  return [head, tail];
};

const list = new SinglyLinkedList()
  .push(3)
  .push(0)
  .push(5)
  .push(2)
  .push(1)
  .push(4);

console.log(rearrangeLinkedList(list, 3));
