package main

import "fmt"

func moveZeroes(nums []int) {
	countZero := 0
	lenNum := len(nums)
	for ind := range nums {
		indReal := ind - countZero
		val := nums[indReal]
		if val == 0 && ind != lenNum-1 {
			indEnd := lenNum - countZero - 1
			for indM := indReal; indM < indEnd; indM++ {
				nums[indM] = nums[indM+1]
			}
			nums[indEnd] = 0
			countZero++
		}
	}
	fmt.Printf("nums:%+v\n", nums)
}

func main() {
	nums := []int{4, 1, 0, 0, 2, 5}
	moveZeroes(nums)
}
