package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1

	for i := 0; i < a-1; i++ {
		list1 = list1.Next
	}

	turnNode := list1

	for i := a - 1; i < b-1; i++ {
		list1 = list1.Next
	}

	turnNode.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = list1

	return head
}

func main() {
	l1 := &ListNode{Val: 10, Next: &ListNode{Val: 1, Next: &ListNode{Val: 13,
		Next: &ListNode{Val: 6, Next: &ListNode{Val: 9, Next: &ListNode{Val: 5,
			Next: nil}}}}}}
	l2 := &ListNode{Val: 1000, Next: &ListNode{Val: 1001,
		Next: &ListNode{Val: 1002, Next: nil}}}

	ln := mergeInBetween(l1, 3, 4, l2)

	fmt.Println(ln)
}
