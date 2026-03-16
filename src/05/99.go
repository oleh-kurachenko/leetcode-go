package main

import "strings"

// findRestaurant implemented in fairly inoptimal way
func findRestaurant(list1 []string, list2 []string) []string {
	result := make([]string, 0)
	indexSum := len(list1) + len(list2)

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if strings.Compare(list1[i], list2[j]) == 0 {
				if i+j < indexSum {
					result = []string{list1[i]}
					indexSum = i + j
				} else if i+j == indexSum {
					result = append(result, list1[i])
				}
			}
		}
	}

	return result
}
