package main

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func addBinary(a string, b string) string {
	s := ""
	var reminder uint8 = 0

	var i int
	var j int
	for i, j = len(a)-1, len(b)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		reminder += a[i] - '0' + b[j] - '0'
		s += string(reminder%2 + '0')
		reminder /= 2
	}
	for ; i >= 0; i-- {
		reminder += a[i] - '0'
		s += string(reminder%2 + '0')
		reminder /= 2
	}
	for ; j >= 0; j-- {
		reminder += b[j] - '0'
		s += string(reminder%2 + '0')
		reminder /= 2
	}

	if reminder != 0 {
		s += string(reminder%2 + '0')
	}

	return reverse(s)
}
