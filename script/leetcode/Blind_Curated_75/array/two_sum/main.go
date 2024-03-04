package main

import "fmt"

func twoSum(num []int, target int) []int {
	hashTmp := map[int]int{}
	for idx, val := range num {
		// find remainder
		rmd := target - val
		idx_rmd, found := hashTmp[rmd]
		if found {
			return []int{idx_rmd, idx}
		}
		// save idx
		hashTmp[val] = idx
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	k := twoSum(nums, target)
	fmt.Printf("k:%+v\n", k)
}
