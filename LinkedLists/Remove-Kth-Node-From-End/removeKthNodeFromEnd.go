package linkedlist

type Node struct {
	Value int
	Next  *Node
}

// Big O: O(n) time | O(1) space, where n is the length of the linked list
func RemoveKthNodeFromEnd(head Node, k int) {
	first, second := &head, &head
	for i := 0; i < k; i++ {
		second = second.Next
	}

	// edge case: the Kth node from end is the head node
	if second == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}

	for second.Next != nil {
		first = first.Next
		second = second.Next
	}

	// now second pointer points to last node in linked list, first pointer
	// points to the previous node of the target node to be removed
	// and the skipped node will be garbage collected
	first.Next = first.Next.Next
}
