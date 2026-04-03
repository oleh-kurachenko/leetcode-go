package main

import "fmt"
import . "leetcode/structs"

import "sort"

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var result []int64

	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		var nextNodes []*TreeNode
		currResult := int64(0)

		for _, node := range nodes {
			currResult += int64(node.Val)
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}

		nodes = nextNodes
		result = append(result, currResult)
	}

	sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })

	if k > len(result) {
		return -1
	}
	return result[k-1]
}

func main() {
	fmt.Println(kthLargestLevelSum(&TreeNode{3, &TreeNode{2,
		&TreeNode{1, nil, nil}, nil}, nil}, 1))
}
