package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hash := map[int]struct{}{}
	for _, val := range nums {
		_, exists := hash[val]
		if exists {
			return true
		}
		hash[val] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	resBool := containsDuplicate(nums)
	fmt.Printf("resBool:%+v\n", resBool)
}
