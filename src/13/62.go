package main

func closestDivisorsOf(num int) []int {
	divisors := []int{1, num}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 && (num/i-i) < divisors[1]-divisors[0] {
			divisors[0] = i
			divisors[1] = num / i
		}
	}

	return divisors
}

func closestDivisors(num int) []int {
	divisors := closestDivisorsOf(num + 1)
	divisors2 := closestDivisorsOf(num + 2)

	if divisors2[1]-divisors2[0] < divisors[1]-divisors[0] {
		return divisors2
	}
	return divisors
}
