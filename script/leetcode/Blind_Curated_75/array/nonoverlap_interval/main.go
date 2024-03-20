package main

import (
	"fmt"
	"slices"
)

func solution(intvs [][]int) int {
	slices.SortFunc(intvs, func(a, b []int) int { return a[1] - b[1] })
	minRemoveCount := 0
	for i := 1; i < len(intvs); i++ {
		if intvs[i][0] < intvs[i-1][1] {
			minRemoveCount++
			intvs[i][1] = intvs[i-1][1]
		}
	}
	return minRemoveCount
}

func main() {
	s := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
