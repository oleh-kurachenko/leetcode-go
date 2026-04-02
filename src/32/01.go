package main

func maximumLengthWithModulo(nums []int, modulo int) int {
	previous := nums[0]
	discardedIdx := -1
	currentCount := 1
	for i := 1; i < len(nums); i++ {
		if (previous+nums[i])%2 == modulo {
			previous = nums[i]
			currentCount++
		} else if discardedIdx == -1 {
			discardedIdx = i
		}
	}
	result := currentCount

	if discardedIdx != -1 {
		previous = nums[discardedIdx]
		currentCount = 1
		for i := discardedIdx + 1; i < len(nums); i++ {
			if (previous+nums[i])%2 == modulo {
				previous = nums[i]
				currentCount++
			}
		}
		result = max(result, currentCount)
	}

	return result
}

func maximumLength(nums []int) int {
	return max(maximumLengthWithModulo(nums, 0),
		maximumLengthWithModulo(nums, 1))
}
