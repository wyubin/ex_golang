package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	indHead := 0
	indTail := len(s) - 1
	for indHead < indTail {
		for !validByte(s[indHead]) && indHead < indTail {
			indHead++
			fmt.Printf("skip to head index:%+v\n", indHead)
		}
		for !validByte(s[indTail]) && indHead < indTail {
			indTail--
			fmt.Printf("skip to tail index:%+v\n", indTail)
		}
		if indHead >= indTail {
			break
		}
		if unicode.ToLower(rune(s[indHead])) != unicode.ToLower(rune(s[indTail])) {
			return false
		}
		indHead++
		indTail--
	}
	return true
}

func validByte(b byte) bool {
	if ('a' <= b && b <= 'z') ||
		('A' <= b && b <= 'Z') ||
		('0' <= b && b <= '9') {
		return true
	}
	return false
}

func main() {
	nums := ",mnnm...."
	k := isPalindrome(nums)
	fmt.Printf("k:%+v\n", k)
}
