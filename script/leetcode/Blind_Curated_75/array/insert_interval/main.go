package main

import (
	"fmt"
	"slices"
)

func insertIntv(intervals [][]int, newInterval []int) [][]int {
	newIntv := append([][]int{newInterval}, intervals...)
	slices.SortFunc(newIntv, func(a, b []int) int { return a[0] - b[0] })
	res := [][]int{newIntv[0]}
	for _, interval := range newIntv[1:] {
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
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	res := insertIntv(intervals, newInterval)
	fmt.Printf("res:%+v\n", res)
}
