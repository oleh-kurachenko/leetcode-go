package main

import (
	"fmt"
	"slices"
	"strings"
)

type replacementOperation struct {
	index  int
	source string
	target string
}

func findReplaceString(s string, indices []int, sources []string,
	targets []string) string {

	replacements := make([]replacementOperation, len(indices))
	for i := range indices {
		replacements[i] = replacementOperation{indices[i], sources[i],
			targets[i]}
	}

	slices.SortFunc(replacements, func(v1 replacementOperation,
		v2 replacementOperation) int {
		if v1.index < v2.index {
			return -1
		}
		if v1.index > v2.index {
			return 1
		}
		return 0
	})

	nextReplacement := 0
	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {
		for ; nextReplacement < len(
			replacements) && replacements[nextReplacement].
			index < i; nextReplacement++ {
		}

		if nextReplacement == len(replacements) {
			sb.WriteString(s[i:])
			break
		}

		replacement := replacements[nextReplacement]

		if i < replacement.index {
			sb.WriteString(s[i:replacement.index])
			i = replacement.index - 1
			continue
		}
		if i+len(replacement.source) <= len(s) &&
			s[i:i+len(replacement.source)] == replacement.source {

			sb.WriteString(replacement.target)
			i += len(replacement.source) - 1
		} else {
			i--
		}
		nextReplacement++
	}

	return sb.String()
}

func main() {
	fmt.Print(findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"},
		[]string{"eee", "ffff"}))
}
