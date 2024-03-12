package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	strLower := strings.ToLower(s)
	left, right := 0, len(strLower)-1
	for left < right {
		for left < right && !isAlphaNumeric(strLower[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(strLower[right]) {
			right--
		}
		if strLower[left] != strLower[right] {
			fmt.Printf("left:%+v right:%+v\n", string(strLower[left]), string(strLower[right]))
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func main() {
	s := "A man, a plan, a canal: Panama"
	k := isPalindrome(s)
	fmt.Printf("k:%+v\n", k)
}
