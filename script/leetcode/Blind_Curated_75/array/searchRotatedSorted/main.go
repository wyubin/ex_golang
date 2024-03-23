package main

import (
	"fmt"
)

// init idxStart, idxEnd
func solution(nums []int, target int) int {
	idxStart, idxEnd := 0, len(nums)-1
	for idxStart <= idxEnd {
		idxMid := (idxStart + idxEnd) / 2
		valMid := nums[idxMid]
		if valMid == target {
			return idxMid
		}
		valStart, valEnd := nums[idxStart], nums[idxEnd]
		if valEnd > valMid {
			if target > valMid && target < valEnd {
				idxStart = idxMid + 1
			} else {
				idxEnd = idxMid - 1
			}
		} else {
			if target < valMid && target >= valStart {
				idxEnd = idxMid - 1
			} else {
				idxStart = idxMid + 1
			}
		}
	}
	return -1
}
func main() {
	s := []int{4, 5, 6, 7, 0, 1, 2}
	k := solution(s, 0)
	fmt.Printf("k: %+v\n", k)
}
