package main

import "sort"

func searchInsert(nums []int, target int) int {
	index := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })

	return index
}
