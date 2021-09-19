package linkedlistpalindrome

type Node struct {
	Value int
	next  *Node
}

func (n *Node) Add(value int) *Node {
	n.next = &Node{Value: value}
	return n.next
}

// Solution 1, recursive approach, since recursive method puts everything on the callstack,
// at the base case we know both the head and tail of the singly linked list without the
// need to use any extra data structure
// Big O: O(n) time | O(n) space, due to the recursive nature, there's n recursive calls on the callstack
// func LinkedListPalindrome(head *Node) bool {
// 	outerNodesAreEqual, _ := isPalindrome(head, head)
// 	return outerNodesAreEqual
// }

// func isPalindrome(leftNode, rightNode *Node) (bool, *Node) {
// 	// base case
// 	if rightNode == nil {
// 		return true, leftNode
// 	}

// 	outerNodesAreEqual, leftNodeToCompare := isPalindrome(leftNode, rightNode.next)
// 	recursiveIsEqual := leftNodeToCompare.Value == rightNode.Value && outerNodesAreEqual
// 	nextMatchingLeftNode := leftNodeToCompare.next

// 	return recursiveIsEqual, nextMatchingLeftNode
// }

// ===================================================================================
// Solution 2, slow/fast pointers iteration approach to find the head of second half of the
// linked list, reverse the second half and get the new head of it (literally the tail of
// the original linked list); iterate from the heads of the first/second half simultaneously
// and compare the two values at each iteration
// Big O: O(n) space | O(1) space
func LinkedListPalindrome(head *Node) bool {
	secondHalfHead := findSecondHalfHead(head)
	reversedSecondHalfHead := reverseLinkedList(secondHalfHead)
	for first, second := head, reversedSecondHalfHead; second != nil; {
		if first.Value != second.Value {
			return false
		}
		first = first.next
		second = second.next
	}
	return true
}

func findSecondHalfHead(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	if fast != nil {
		// the length of linked list is odd
		return slow.next
	}
	// the length of linked list is even
	return slow
}

func reverseLinkedList(head *Node) *Node {
	var p1, p2, p3 *Node
	p2 = head
	for p2 != nil {
		p3 = p2.next
		p2.next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
