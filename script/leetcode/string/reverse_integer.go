package main

import (
	"fmt"
)

func reverse(x int) int {
	indPos := true
	if x < 0 {
		indPos = false
		x = 0 - x
	} else if x == 0 {
		return x
	}
	arrInt := []int{}
	for x >= 1 {
		ras := x % 10
		x = x / 10
		arrInt = append([]int{ras}, arrInt...)
	}
	y := 0
	fold := 1
	for _, val := range arrInt {
		y += val * fold
		fold *= 10
	}
	if y > (1 << 31) {
		return 0
	}
	if !indPos {
		y = 0 - y
	}
	return y
}

func main() {
	nums := -153
	k := reverse(nums)
	fmt.Printf("k:%+v\n", k)
}
