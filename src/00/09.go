package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var x_copy int64 = int64(x)
	var x_reverse int64 = 0

	for ; x > 0; x /= 10 {
		x_reverse = x_reverse*10 + int64(x%10)
	}

	return x_copy == x_reverse
}

func main() {
	fmt.Printf("%v", isPalindrome(121))
}
