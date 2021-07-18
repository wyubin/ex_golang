package main

import "fmt"

func singleNumber(nums []int) int {
	hash := map[int]bool{}
	for _, val := range nums {
		indDup, exists := hash[val]
		if !exists {
			hash[val] = true
		} else if indDup {
			delete(hash, val)
		}
	}
	var res int
	for val := range hash {
		res = val
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	k := singleNumber(nums)
	fmt.Printf("k:%+v\n", k)
}
