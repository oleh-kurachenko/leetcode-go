package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var inSubstring [256]bool
	maxLength := 0
	tailPosition := 0
	for i, c := range s {
		if inSubstring[c] {
			for ; s[tailPosition] != byte(c); tailPosition++ {
				inSubstring[s[tailPosition]] = false
			}
			tailPosition++
		} else {
			inSubstring[c] = true
			maxLength = max(maxLength, i-tailPosition+1)
		}
	}

	return maxLength
}

func main() {
	fmt.Printf("%v", lengthOfLongestSubstring("abcabcbb"))
}
