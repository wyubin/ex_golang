package main

import "fmt"

func firstUniqChar(s string) int {
	indLast := map[byte]int{}
	for _, val := range []byte(s) {
		indLast[val] += 1
	}
	for ind, val := range []byte(s) {
		if indLast[val] == 1 {
			return ind
		}
	}
	return -1
}

func main() {
	nums := "aadadaad"
	k := firstUniqChar(nums)
	fmt.Printf("k:%d\n", k)
}
