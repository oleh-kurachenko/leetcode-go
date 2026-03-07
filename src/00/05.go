package main

import "fmt"

func isPalindromeString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	max_length := 1
	max_string := s[0:1]

	for i := 0; i < len(s)-max_length; i++ {
		for j := i + 1 + max_length; j <= len(s); j++ {
			if isPalindromeString(s[i:j]) {
				max_length = j - i
				max_string = s[i:j]
			}
		}
	}

	return max_string
}

func main() {
	fmt.Printf("%v", longestPalindrome("bb"))
}
