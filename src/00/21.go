package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root *ListNode = nil
	var tail **ListNode = &root

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			*tail = list1
			tail = &list1.Next
			list1 = list1.Next
		} else {
			*tail = list2
			tail = &list2.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		*tail = list1
	}
	if list2 != nil {
		*tail = list2
	}

	return root
}
