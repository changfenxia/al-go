package main

import "fmt"

func main() {
	test := []int{3, 2, 3}
	fmt.Print(getMajority(test))
}

func getMajority(list []int) int {
	var res = 0
	var count = 0

	for i := 0; i < len(list); i += 1 {
		if count == 0 {
			res = list[i]
		}
		if res == list[i] {
			count += 1
		} else {
			count -= 1
		}
	}

	return res
}
