package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func main() {
	var nums = []int{6, 2, 7, 11, 15}
	var target = 21

	fmt.Print(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i += 1 {
		cur := nums[i]
		if val, ok := m[cur]; ok {
			return []int{val, i}
		}
		remainder := target - cur
		m[remainder] = i
	}

	return []int{}
}
