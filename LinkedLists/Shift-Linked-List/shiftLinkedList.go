// Given a linked list and an arbitrary integer k, shift the node elements with an offset equals to k
// k > 0 shifting forward (removing from tail and link back to head); k < 0 shifting backward

package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func main() {
	srcList := LinkedList{head: nil, tail: nil}
	srcList.push(0)
	srcList.push(1)
	srcList.push(2)
	srcList.push(3)
	srcList.push(4)
	srcList.push(5)

	shiftLinkedList(&srcList, -2)
	for i := srcList.head; i != nil; i = i.next {
		fmt.Printf("%+v\t", i.value)
	}
}

// Tail definition here only for ease of pushing (initializing the src list), this question limits that we have no reference to the tail node
func (list *LinkedList) push(val int) *LinkedList {
	newNode := &Node{
		value: val,
		next:  nil,
	}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
	return list
}

// Big O: O(n) time, traversing n+(n-k) nodes | O(1) space
func shiftLinkedList(list *LinkedList, k int) *LinkedList {
	var oldTail *Node
	length := 0
	for i := list.head; i != nil; i = i.next {
		length += 1
		oldTail = i
	}

	// make offset an positive value, only differece is for negative k, offset would be counted from head, not tail
	offset := int(math.Abs(float64(k))) % length
	if offset == 0 {
		return list
	}

	var newTail *Node
	// newTailPosition = length - offset if (0 < offset < length); if -length < offset < 0 then newTailPosition = offset
	var newTailPosition = length - offset
	if k < 0 {
		newTailPosition = offset
	}
	for i, counter := list.head, 0; counter < newTailPosition; i, counter = i.next, counter+1 {
		newTail = i
	}

	newHead := newTail.next
	oldTail.next = list.head
	list.head = newHead
	list.tail = newTail
	newTail.next = nil
	return list
}
