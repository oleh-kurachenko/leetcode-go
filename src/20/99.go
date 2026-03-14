package main

import "fmt"

import "slices"

func maxSubsequence(nums []int, k int) []int {
	numsIdx := make([]int, len(nums))
	for i := 0; i < len(numsIdx); i++ {
		numsIdx[i] = i
	}

	slices.SortFunc(numsIdx, func(i, j int) int {
		if nums[i] < nums[j] {
			return -1
		} else if nums[i] > nums[j] {
			return 1
		}
		return 0
	})

	numsPositionIdx := make([]int, len(nums))
	for i := 0; i < len(numsPositionIdx); i++ {
		numsPositionIdx[numsIdx[i]] = i
	}

	result := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		if numsPositionIdx[i] > len(nums)-k-1 {
			result = append(result, nums[i])
		}
	}

	return result
}

func main() {
	fmt.Print(maxSubsequence([]int{63, -74, 61, -17, -55, -59, -10, 2, -60,
		-65}, 9))
}
