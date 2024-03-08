package main

import "fmt"

func canJump(nums []int) bool {
	maxIdx := 1
	lastIdx := len(nums) - 1
	for idx, num := range nums {
		maxIdx--
		if num > maxIdx {
			maxIdx = num
		}
		if maxIdx == 0 && idx != lastIdx {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	k := canJump(nums)
	fmt.Printf("k:%+v\n", k)
}
