package main

import "fmt"

func plusOne(digits []int) []int {
	lenNum := len(digits)
	getTen := true
	resDigit := []int{}
	for ind := lenNum - 1; ind >= 0; ind-- {
		digitCurr := digits[ind]
		if getTen {
			digitCurr++
			getTen = digitCurr == 10
		}
		resDigit = append([]int{digitCurr % 10}, resDigit...)
	}
	if getTen {
		resDigit = append([]int{1}, resDigit...)
	}
	return resDigit
}

func main() {
	nums := []int{9, 9}
	k := plusOne(nums)
	fmt.Printf("k:%+v\n", k)
}
