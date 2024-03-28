package main

import (
	"fmt"
	"slices"
)

func solution(srcStr string, chTime int) int {
	// init char2Count
	char2Count := make([]int, 26)
	idxStart, idxEnd := 0, 0
	maxRepeat, currRepeat := 0, 0
	for idxEnd < len(srcStr) {
		charEnd := int(rune(srcStr[idxEnd]) - 'A')
		char2Count[charEnd]++
		currRepeat = slices.Max(char2Count)
		currLen := idxEnd - idxStart + 1
		idxEnd++
		if currLen-currRepeat <= chTime {
			maxRepeat = slices.Max([]int{maxRepeat, currLen})
			continue
		}
		// move idxStart
		charStart := int(rune(srcStr[idxStart]) - 'A')
		char2Count[charStart]--
		idxStart++
	}
	return maxRepeat
}

func main() {
	s := "AABABBA"
	n := 1
	k := solution(s, n)
	fmt.Printf("k: %+v\n", k)
}
