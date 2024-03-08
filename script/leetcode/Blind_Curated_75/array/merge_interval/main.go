package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	res := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		resLen := len(res)
		if res[resLen-1][1] < interval[0] {
			res = append(res, interval)
		} else if interval[1] > res[resLen-1][1] {
			res[len(res)-1][1] = interval[1]
		}
	}
	return res
}

func main() {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	k := merge(nums)
	fmt.Printf("k:%+v\n", k)
}
