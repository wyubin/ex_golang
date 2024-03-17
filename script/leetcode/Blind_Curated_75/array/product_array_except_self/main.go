package main

import "fmt"

func solution(nums []int) []int {
	numsLen := len(nums)
	res := make([]int, numsLen)
	prefix := make([]int, numsLen)
	suffix := make([]int, numsLen)
	prefix[0] = 1
	for i := 1; i < numsLen; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	suffix[numsLen-1] = 1
	for i := numsLen - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i := 0; i < numsLen; i++ {
		res[i] = prefix[i] * suffix[i]
	}
	return res
}
func main() {
	s := []int{1, 2, 3, 4}
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
