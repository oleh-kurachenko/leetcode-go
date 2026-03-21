package main

func resultsArray(nums []int, k int) []int {
	results := make([]int, 0, len(nums)-k+1)

	consecutive := 0
	for i, val := range nums {
		if i == 0 || nums[i-1]+1 == val {
			consecutive++
		} else {
			consecutive = 1
		}

		if k <= i+1 {
			if consecutive >= k {
				results = append(results, val)
			} else {
				results = append(results, -1)
			}
		}
	}

	return results
}
