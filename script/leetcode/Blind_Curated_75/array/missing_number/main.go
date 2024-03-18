package main

import (
	"fmt"
	"slices"
)

func solution(nums []int) int {
	lenNums := len(nums)
	slices.Sort(nums)
	var idx int
	for idx = 0; idx < lenNums; idx++ {
		if nums[idx] != idx {
			return idx
		}
	}
	return lenNums
}

func main() {
	s := []int{3, 0, 1}
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
