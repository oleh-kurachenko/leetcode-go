package main

import "strings"

func lengthOfLastWord(s string) int {
	word_end := strings.LastIndexFunc(s, func(r rune) bool { return r != ' ' })
	word_start := strings.LastIndex(s[:word_end], " ")

	return word_end - word_start
}
