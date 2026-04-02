package main

import "fmt"

func goodIndices(nums []int, k int) []int {
	postNondecr := make([]int, len(nums))

	postNondecr[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			postNondecr[i] = postNondecr[i+1] + 1
		} else {
			postNondecr[i] = 1
		}
	}

	preNonincr := 1
	var result []int
	for i := 1; i < len(nums)-k; i++ {
		if i >= k && preNonincr >= k && postNondecr[i+1] >= k {
			result = append(result, i)
		}

		if nums[i] <= nums[i-1] {
			preNonincr++
		} else {
			preNonincr = 1
		}
	}

	return result
}

func main() {
	fmt.Println(goodIndices([]int{878724, 201541, 179099, 98437, 35765, 327555,
		475851, 598885, 849470, 943442}, 4))
}
