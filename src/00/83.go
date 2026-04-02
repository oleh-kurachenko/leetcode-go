package main

func deleteDuplicates(head *ListNode) *ListNode {
	root := head

	prev := head
	for current := head; current != nil; current = current.Next {
		if current.Val == prev.Val {
			prev.Next = current.Next
		} else {
			prev = current
		}
	}

	return root
}
