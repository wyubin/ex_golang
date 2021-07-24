package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	// 最短先放在前面
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	lenStr := len(strs)
	if lenStr == 1 {
		return strs[0]
	}
	var indCount int
	for indCompare := range strs[0] {
		for indStr := 1; indStr < len(strs); indStr++ {
			if strs[0][indCompare] != strs[indStr][indCompare] {
				return strs[0][:indCompare]
			}
		}
		indCount = indCompare + 1
	}
	return strs[0][:indCount]
}

func main() {
	nums := []string{"flower", "flow", "flight"}
	k := longestCommonPrefix(nums)
	fmt.Printf("k:%s\n", k)
}
