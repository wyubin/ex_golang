package main

import (
	"fmt"
	"slices"
)

func getMaxProduct(nums []int) int {
	maxProduct := 0
	var currProduct int

	currProduct = 1
	for i := 0; i < len(nums); i++ {
		currProduct = currProduct * nums[i]
		maxProduct = slices.Max([]int{currProduct, maxProduct})
		if currProduct == 0 {
			currProduct = 1
		}
	}
	// reverse
	currProduct = 1
	for i := len(nums) - 1; i >= 0; i-- {
		currProduct = currProduct * nums[i]
		maxProduct = slices.Max([]int{currProduct, maxProduct})
		if currProduct == 0 {
			currProduct = 1
		}
	}
	return maxProduct
}

func main() {
	s := []int{2, 3, -2, 4}
	k := getMaxProduct(s)
	fmt.Printf("k: %+v\n", k)
}
