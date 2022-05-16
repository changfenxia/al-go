package main

import "fmt"

func main() {
	test := []int{2, 2, 2, 2, 3, 3, 2, 3, 2, 5, 3, 3}
	fmt.Print(getMajority(test))
}

func getMajority(list []int) int {
	var res = list[0]
	var count = 1

	for i := 1; i < len(list); i += 1 {
		if res == list[i] {
			count += 1
		} else {
			count -= 1
		}

		if count < 0 {
			res = list[i]
		}
	}

	return res
}
