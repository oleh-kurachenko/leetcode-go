package main

func countMajoritySubarrays(nums []int, target int) int {
	if nums[0] == target {
		nums[0] = 1
	} else {
		nums[0] = 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == target {
			nums[i] = nums[i-1] + 1
		} else {
			nums[i] = nums[i-1]
		}
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			target_count := nums[j]
			if i > 0 {
				target_count -= nums[i-1]
			}

			if target_count > (j-i+1)/2 {
				result++
			}
		}
	}

	return result
}
