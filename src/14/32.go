package main

import "fmt"

func maxDiff(num int) int {
	pow := 1

	for ; num/pow > 9; pow *= 10 {
	}

	isFirstDigit := true
	maximizeReplace := -1
	minimizeReplace := -1
	minimizeValue := -1
	a := 0
	b := 0
	for ; pow > 0; pow /= 10 {
		currentDigit := num / pow

		if maximizeReplace == -1 && currentDigit != 9 {
			maximizeReplace = currentDigit
		}

		if isFirstDigit {
			if currentDigit != 1 {
				minimizeReplace = currentDigit
				minimizeValue = 1
			}
		} else {
			if minimizeReplace == -1 && currentDigit != 0 && currentDigit != 1 {
				minimizeReplace = currentDigit
				minimizeValue = 0
			}
		}

		if currentDigit == maximizeReplace {
			a += pow * 9
		} else {
			a += (currentDigit) * pow
		}

		if currentDigit == minimizeReplace {
			b += pow * minimizeValue
		} else {
			b += (currentDigit) * pow
		}

		num = num % pow
		isFirstDigit = false
	}

	return a - b
}

func main() {
	fmt.Println(maxDiff(555))
}
