package easyother

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for ind, val := range nums2 {
		nums1[m+ind] = val
	}
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })

}
