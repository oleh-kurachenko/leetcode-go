package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := l1
	reminder := 0
	prev_l1 := l1

	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		reminder += l1.Val + l2.Val
		l1.Val = reminder % 10
		reminder /= 10
		prev_l1 = l1
	}
	if l1 == nil {
		prev_l1.Next = l2
		l1 = l2
	}

	for ; reminder > 0 && l1 != nil; l1 = l1.Next {
		reminder += l1.Val
		l1.Val = reminder % 10
		reminder /= 10
		prev_l1 = l1
	}
	if reminder > 0 {
		prev_l1.Next = &ListNode{Val: reminder}
	}

	return root
}
