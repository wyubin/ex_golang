package main

import "fmt"

func maxSubarray(nums []int) int {
	var max int
	sumSubarr := 0
	for _, num := range nums {
		sumSubarr += num
		if sumSubarr > max {
			max = sumSubarr
		}
		if sumSubarr < 0 {
			sumSubarr = 0
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	k := maxSubarray(nums)
	fmt.Printf("k:%+v\n", k)
}
