package main

func maximumUniqueSubarray(nums []int) int {
	isUsed := make([]bool, 10001)

	frontPtr := 0
	sum := nums[0]
	maxSum := sum
	isUsed[nums[0]] = true

	for i := 1; i < len(nums); i++ {
		if isUsed[nums[i]] {
			for ; nums[frontPtr] != nums[i]; frontPtr++ {
				isUsed[nums[frontPtr]] = false
				sum -= nums[frontPtr]
			}
			sum -= nums[frontPtr]
			frontPtr++
		}

		sum += nums[i]
		isUsed[nums[i]] = true

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
