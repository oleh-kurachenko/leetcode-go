package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	lPtr := head
	lLen := 1

	for ; lPtr.Next != nil; lPtr = lPtr.Next {
		lLen++
	}

	k %= lLen

	lPtr.Next = head
	for i := 0; i < lLen-k; i++ {
		lPtr = lPtr.Next
	}

	lReturn := lPtr.Next
	lPtr.Next = nil

	return lReturn
}
