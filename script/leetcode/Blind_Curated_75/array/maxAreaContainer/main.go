package main

import (
	"fmt"
	"slices"
)

func maxArea(heights []int) int {
	maxA := 0
	idxLeft := 0
	idxRight := len(heights) - 1

	for idxLeft < idxRight {
		area := slices.Min([]int{heights[idxLeft], heights[idxRight]}) * (idxRight - idxLeft)
		if area > maxA {
			maxA = area
		}
		if heights[idxLeft] < heights[idxRight] {
			idxLeft++
		} else {
			idxRight--
		}
	}
	return maxA
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	k := maxArea(heights)
	fmt.Printf("k:%+v\n", k)
}
