package main

import (
	"fmt"
	"slices"
)

func maxArea(heights []int) int {
	maxA := 0

	for i := 0; i < len(heights); i++ {
		for j := len(heights) - 1; j > i; j-- {
			area := slices.Min([]int{heights[i], heights[j]}) * (j - i)
			if area > maxA {
				maxA = area
			}
		}
	}
	return maxA
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	k := maxArea(heights)
	fmt.Printf("k:%+v\n", k)
}
