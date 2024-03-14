package main

import (
	"fmt"
	"slices"
)

func maxRob(nums []int, idx int, dp map[int]int) int {
	var maxMoney int
	if idx < 0 {
		return 0
	}
	maxMoney, found := dp[idx]
	if found {
		return maxMoney
	}
	maxMoney = slices.Max([]int{nums[idx] + maxRob(nums, idx-2, dp), maxRob(nums, idx-1, dp)})
	dp[idx] = maxMoney
	return maxMoney
}

func solution(nums []int) int {
	dp := make(map[int]int, len(nums))
	return maxRob(nums, len(nums)-1, dp)
}

func main() {
	s := []int{2, 7, 9, 3, 1}
	k := solution(s)
	fmt.Printf("k: %d\n", k)
}
