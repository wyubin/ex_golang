package main

import "fmt"

func solution(nums []int) int {
	startIdx, endIdx := 0, len(nums)-1
	var midIdx int
	for startIdx < endIdx {
		midIdx = (startIdx + endIdx) / 2
		if nums[midIdx] > nums[endIdx] {
			startIdx = midIdx + 1
		} else {
			endIdx = midIdx
		}
	}
	return nums[startIdx]
}

func main() {
	s := []int{3, 4, 5, 1, 2}
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
