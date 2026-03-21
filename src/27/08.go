package main

import "slices"

func maxStrength(nums []int) (result int64) {
	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if i+1 < len(nums) && nums[i+1] < 0 {
				if result == 0 {
					result = 1
				}
				result *= int64(nums[i]) * int64(nums[i+1])
				i++
			}
		} else if nums[i] > 0 {
			if result == 0 {
				result = 1
			}
			result *= int64(nums[i])
		}
	}

	if result == 0 {
		result = int64(nums[len(nums)-1])
	}

	return
}
