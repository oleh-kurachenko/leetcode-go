package main

import . "leetcode/structs"

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(inorderTraversal(root.Left), root.Val),
		inorderTraversal(root.Right)...)
}
