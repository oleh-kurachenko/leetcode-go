package main

func findFirstNonEqual(str1 string, str2 string) int {
	i := 0
	for ; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			return i
		}
	}
	return i
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = prefix[0:findFirstNonEqual(prefix, strs[i])]
	}

	return prefix
}
