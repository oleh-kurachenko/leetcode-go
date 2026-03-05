package main

func removeElement(nums []int, val int) int {
	l := 0
	for _, v := range nums {
		if v != val {
			nums[l] = v
			l++
		}
	}

	return l
}

func main() {
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}
