package main

import (
	"fmt"
	"strings"
)

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func solution(encStr string) bool {
	encStrLower := strings.ToLower(encStr)
	idxStart, idxEnd := 0, len(encStr)-1
	for idxStart < idxEnd {
		for !isAlphaNumeric(encStrLower[idxStart]) {
			idxStart++
		}
		for !isAlphaNumeric(encStrLower[idxEnd]) {
			idxEnd--
		}
		if encStrLower[idxStart] != encStrLower[idxEnd] {
			return false
		}
		idxStart++
		idxEnd--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
