// golang entry for each AlgoExpert challenge
// #[Linked List]/Linked List Palindrome

package main

import (
	"fmt"

	linkedList "github.com/letempest/algoexpert-in-action/LinkedLists/Linked-List-Palindrome"
)

func main() {
	// example 1, should yield true
	head := linkedList.Node{Value: 0}
	func(node *linkedList.Node, vals ...int) {
		nextNode := node
		for _, num := range vals {
			nextNode = nextNode.Add(num)
		}
	}(&head, 1, 2, 2, 1, 0)

	fmt.Printf("Is palindrome? %v\n", linkedList.LinkedListPalindrome(&head))

	// example 2, should yield false
	head = linkedList.Node{Value: 0}
	func(node *linkedList.Node, vals ...int) {
		nextNode := node
		for _, num := range vals {
			nextNode = nextNode.Add(num)
		}
	}(&head, 1, 2, 2, 1, 1)

	fmt.Printf("Is palindrome? %v\n", linkedList.LinkedListPalindrome(&head))
}
