package main

import (
	"fmt"
	"slices"
)

func solution(inStr string, lenWild int) int {
	countChar := make([]int, 26)
	maxRepeat := 0
	resLen := 0
	idxStart, idxEnd := 0, 0
	var intStart, intEnd int
	// move window
	for idxEnd < len(inStr) {
		intEnd = int(rune(inStr[idxEnd] - 'A'))
		countChar[intEnd]++
		maxRepeat = slices.Max([]int{maxRepeat, countChar[intEnd]})
		if idxEnd-idxStart+1-maxRepeat > lenWild {
			intStart = int(rune(inStr[idxStart] - 'A'))
			countChar[intStart]--
			idxStart++
		}
		// record total len
		resLen = slices.Max([]int{resLen, idxEnd - idxStart + 1})
		idxEnd++
	}
	return resLen
}

func main() {
	s := "AABABBA"
	n := 1
	k := solution(s, n)
	fmt.Printf("k: %+v\n", k)
}
