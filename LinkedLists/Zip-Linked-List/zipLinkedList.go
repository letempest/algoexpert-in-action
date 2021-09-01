package main

import "fmt"

type LinkedList struct {
	value int
	next  *LinkedList
}

func (n *LinkedList) add(value int) *LinkedList {
	n.next = &LinkedList{value: value}
	return n.next
}

func main() {
	linkedList := LinkedList{value: 1}
	linkedList.add(2).add(3).add(4).add(5).add(6)

	zipLinkedList(&linkedList)
	for node := linkedList; true; {
		fmt.Printf("%+v\t", node.value)
		if node.next == nil {
			break
		}
		node = *node.next
	}
}

// can divide the task into three parts, 1) split the original linked list roughly into two half and get reference of both heads (don't forget to set firsthalf.tail.next = nil);
// 2) reverse the second half; 3) interweave the first half with the reversed second half, and voila, the linked list is zipped
// Big O: O(n) time | O(1) space
func zipLinkedList(linkedList *LinkedList) *LinkedList {
	// if the source linked list is of lenth 1 or 2, then don't need to do anything, though the implementation below can handle these edge cases too
	if linkedList.next == nil || linkedList.next.next == nil {
		return linkedList
	}

	firstHalfHead := linkedList
	secondHalfHead := splitLinkedList(linkedList)

	reversedSecondHalfHead := reverseLinkedList(secondHalfHead)

	return interweaveLinkedLists(firstHalfHead, reversedSecondHalfHead)
}

func splitLinkedList(linkedList *LinkedList) *LinkedList {
	slowIterator, fastIterator := linkedList, linkedList
	for fastIterator != nil && fastIterator.next != nil {
		slowIterator = slowIterator.next
		fastIterator = fastIterator.next.next
	}
	secondHalfHead := slowIterator.next
	slowIterator.next = nil
	return secondHalfHead
}

func reverseLinkedList(linkedList *LinkedList) *LinkedList {
	var previousNode, currentNode *LinkedList
	currentNode = linkedList
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

func interweaveLinkedLists(firstLinkedList, secondLinkedList *LinkedList) *LinkedList {
	firstListIterator, secondListIterator := firstLinkedList, secondLinkedList

	// For this implementation, second half is definitely shorter than first half due to the way we split them, so it's also okay to just check the second pointer
	for firstListIterator != nil && secondListIterator != nil {
		firstListIteratorNext, secondListIteratorNext := firstListIterator.next, secondListIterator.next

		firstListIterator.next = secondListIterator
		secondListIterator.next = firstListIteratorNext

		firstListIterator = firstListIteratorNext
		secondListIterator = secondListIteratorNext
	}

	return firstLinkedList
}
