package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	name_pos := 0
	typed_pos := 0

	for typed_pos < len(typed) {
		if name_pos < len(name) && name[name_pos] == typed[typed_pos] {
			name_pos++
			typed_pos++
		} else if name_pos > 0 && name[name_pos-1] == typed[typed_pos] {
			typed_pos++
		} else {
			return false
		}
	}

	return name_pos == len(name)
}

func main() {
	fmt.Print(isLongPressedName("alex", "aaleex"))
}
