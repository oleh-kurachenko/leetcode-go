package main

func climbStairs(n int) int {
	prev := 1
	curr := 1

	for i := 1; i < n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}

	return curr
}
