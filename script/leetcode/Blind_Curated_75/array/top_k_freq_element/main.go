package main

import (
	"fmt"
	"slices"
)

func solution(nums []int, topLen int) []int {
	num2count := map[int]int{}
	for _, num := range nums {
		num2count[num]++
	}
	pairs := [][]int{}
	for num, count := range num2count {
		pairs = append(pairs, []int{num, count})
	}
	slices.SortFunc(pairs, func(a, b []int) int { return b[1] - a[1] })
	res := []int{}
	for _, pair := range pairs[:topLen] {
		res = append(res, pair[0])
	}
	return res
}

func main() {
	s := []int{1, 1, 1, 2, 2, 3}
	k := solution(s, 2)
	fmt.Printf("k: %+v\n", k)
}
