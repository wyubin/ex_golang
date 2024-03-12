package main

import "fmt"

func longestConsecutive(nums []int) int {
	num2left := map[int]bool{}
	for _, num := range nums {
		num2left[num] = true
	}
	maxLen := 0
	for i := 0; i < len(nums) && num2left[nums[i]]; i++ {
		num := nums[i]
		delete(num2left, num)
		currLen := 1
		preNum := num - 1
		nextNum := num + 1
		for num2left[preNum] {
			currLen++
			delete(num2left, preNum)
			preNum--
		}
		for num2left[nextNum] {
			currLen++
			delete(num2left, nextNum)
			nextNum++
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

func main() {

	s := []int{100, 4, 200, 1, 3, 2}
	k := longestConsecutive(s)
	fmt.Printf("k:%+v\n", k)
}
