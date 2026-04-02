package main

func countSubstrings(s string) int {
	result := 0

	for i := range s {
		for left_i, right_i := i, i; left_i >= 0 &&
			right_i < len(s) && s[left_i] == s[right_i]; left_i,
			right_i = left_i-1, right_i+1 {

			result++
		}
		for left_i, right_i := i, i+1; left_i >= 0 &&
			right_i < len(s) && s[left_i] == s[right_i]; left_i,
			right_i = left_i-1, right_i+1 {

			result++
		}
	}

	return result
}
