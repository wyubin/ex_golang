package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for ind, val := range nums {
		matchInd, exists := hash[val]
		if exists {
			return []int{matchInd, ind}
		}
		hash[target-val] = ind
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	k := twoSum(nums, 9)
	fmt.Printf("k:%+v\n", k)
}
