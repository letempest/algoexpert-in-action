// Given a sorted! linked list, remove all duplicate node from it
// Input:  1 -> 3 -> 4 -> 4 -> 4 -> 5 -> 6 -> 6
// Output: 1 -> 3 -> 4 -> 5 -> 6
// Big O: O(n) Time | O(1) Space

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

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

func main() {
	list := LinkedList{
		head: nil,
		tail: nil,
	}

	list.push(1)
	list.push(3)
	list.push(4)
	list.push(4)
	list.push(4)
	list.push(5)
	list.push(6)
	list.push(6)

	fmt.Println("-----before-------")
	for i := list.head; i != nil; i = i.next {
		fmt.Printf("%+v\n", i.value)
	}
	fmt.Println("-----after-------")
	list = *removeDuplicatesFromLinkedList(&list)
	for i := list.head; i != nil; i = i.next {
		fmt.Printf("%+v\n", i.value)
	}
}

func removeDuplicatesFromLinkedList(list *LinkedList) *LinkedList {
	currentNode := list.head
	nextDistinctNode := currentNode.next
	for currentNode != nil {
		nextDistinctNode = currentNode.next
		for nextDistinctNode != nil && nextDistinctNode.value == currentNode.value {
			nextDistinctNode = nextDistinctNode.next
		}
		currentNode.next = nextDistinctNode
		currentNode = nextDistinctNode
	}

	return list
}
