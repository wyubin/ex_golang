package main

import "fmt"

func solution(nums []int, target int) int {
	idxLow := 0
	idxHigh := len(nums) - 1
	for idxLow <= idxHigh {
		idxMid := (idxLow + idxHigh) / 2
		if nums[idxMid] == target {
			return idxMid
		}
		if nums[idxLow] <= nums[idxMid] {
			if nums[idxLow] <= target && target < nums[idxMid] {
				idxHigh = idxMid - 1
			} else {
				idxLow = idxMid + 1
			}
		} else {
			if nums[idxMid] < target && target <= nums[idxHigh] {
				idxLow = idxMid + 1
			} else {
				idxHigh = idxMid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	k := solution(nums, target)
	fmt.Printf("k:%+v\n", k)
}
