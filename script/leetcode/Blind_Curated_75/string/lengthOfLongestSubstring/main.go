package main

import (
	"fmt"
	"slices"
)

func solution(inStr string) int {
	maxLen := 0
	char2idx := map[rune]int{}
	idxStart := 0
	for idxCurr, currRune := range inStr {
		preIdx, found := char2idx[currRune]
		if found {
			idxStart = preIdx + 1
		} else {
			maxLen = slices.Max([]int{maxLen, idxCurr - idxStart + 1})
		}
		char2idx[currRune] = idxCurr
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
