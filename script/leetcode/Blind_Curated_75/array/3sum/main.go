package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	currSet := ""
	sets := map[string]struct{}{}
	numsSorted := nums[0:]
	slices.Sort(numsSorted)
	for i := 0; i < len(numsSorted)-2; i++ {
		j := i + 1
		k := len(numsSorted) - 1
		for j < k {
			sum := numsSorted[i] + numsSorted[j] + numsSorted[k]
			if sum == 0 {
				// check if duplicate
				currSet = fmt.Sprintf("%d,%d,%d", numsSorted[i], numsSorted[j], numsSorted[k])
				if _, found := sets[currSet]; !found {
					res = append(res, []int{numsSorted[i], numsSorted[j], numsSorted[k]})
					sets[currSet] = struct{}{}
				}
				j++
				k--
				for j < k && numsSorted[j] == numsSorted[j-1] {
					j++
				}
				for j < k && numsSorted[k] == numsSorted[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	k := threeSum(nums)
	fmt.Printf("k:%+v\n", k)
}
