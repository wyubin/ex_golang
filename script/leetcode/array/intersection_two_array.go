package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	numsTmp := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	ind1, ind2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	for ind1 = 0; ind1 < len1 && ind2 < len2; ind1++ {
		for nums1[ind1] >= nums2[ind2] {
			if nums1[ind1] == nums2[ind2] {
				numsTmp = append(numsTmp, nums1[ind1])
				ind2++
				break
			} else {
				ind2++
				if !(ind2 < len2) {
					break
				}
			}
		}
	}
	return numsTmp
}

func main() {
	nums := []int{4, 1, 2, 1, 2, 5}
	nums2 := []int{4, 1, 2, 1, 9}
	k := intersect(nums, nums2)
	fmt.Printf("k:%+v\n", k)
}
