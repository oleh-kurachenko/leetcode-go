package main

func romanToInt(s string) int {
	var previous byte = '_'
	var result = 0

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if previous != 'M' && previous != 'D' {
				result += 100
			} else {
				result -= 100
			}
		case 'L':
			result += 50
		case 'X':
			if previous != 'L' && previous != 'C' {
				result += 10
			} else {
				result -= 10
			}
		case 'V':
			result += 5
		case 'I':
			if previous != 'V' && previous != 'X' {
				result++
			} else {
				result--
			}
		}
		previous = s[i]
	}

	return result
}
