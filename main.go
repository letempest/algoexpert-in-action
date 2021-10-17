// golang entry for each AlgoExpert challenge
// #[Linked List]/Remove Kth Node From End

package main

import (
	"fmt"

	linkedlist "github.com/letempest/algoexpert-in-action/LinkedLists/Remove-Kth-Node-From-End"
)

func main() {
	head := linkedlist.Node{Value: 0}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for currentNode, i := &head, 0; i < len(values); i++ {
		nextNode := &linkedlist.Node{Value: values[i]}
		currentNode.Next = nextNode
		currentNode = nextNode
	}

	linkedlist.RemoveKthNodeFromEnd(head, 4)

	for currentNode := &head; currentNode != nil; {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}
