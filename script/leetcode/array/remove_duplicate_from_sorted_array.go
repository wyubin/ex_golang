package main

import "fmt"

func removeDuplicates(nums []int) int {
	ind := 0
	var curVal int
	for _, val := range nums {
		if ind == 0 || curVal != val {
			curVal = val
			nums[ind] = curVal
			ind++
		}
	}
	nums = nums[:ind]
	return ind
}

func main() {
	nums := []int{-1, 0, 0, 0, 0, 3, 3}
	k := removeDuplicates(nums)
	fmt.Printf("nums:%+v\n", nums)
	fmt.Printf("k:%+v\n", k)
}
