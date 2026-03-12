package main

func areSymmetric(rootA *TreeNode, rootB *TreeNode) bool {
	if rootA == nil && rootB == nil {
		return true
	}

	if rootA == nil || rootB == nil || rootA.Val != rootB.Val {
		return false
	}

	return areSymmetric(rootA.Left, rootB.Right) &&
		areSymmetric(rootA.Right, rootB.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return areSymmetric(root.Left, root.Right)
}
