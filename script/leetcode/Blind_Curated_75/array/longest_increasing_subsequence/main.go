package main

import (
	"fmt"
	"slices"
)

func solution(nums []int) int {
	numLen := len(nums)
	// init dp
	dp := make([]int, numLen)
	for idx := range dp {
		dp[idx] = 1
	}
	// top->bottom and record max in dp, caculate all possible combination
	for i := 0; i < numLen; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = slices.Max([]int{dp[i], dp[j] + 1})
			}
		}
	}
	return slices.Max(dp)
}

func main() {
	s := []int{10, 9, 2, 5, 3, 7, 101, 18}
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
