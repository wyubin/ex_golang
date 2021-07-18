package main

import "fmt"

func rotate(nums []int, k int) {
	lenNums := len(nums)
	stepOp := k % lenNums
	indLast := lenNums - stepOp
	fmt.Printf("indLast:%d\n", indLast)
	numsTmp := append(nums[indLast:], nums[:indLast]...)
	for ind, val := range numsTmp {
		nums[ind] = val
	}
	fmt.Printf("nums:%+v\n", nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
}
