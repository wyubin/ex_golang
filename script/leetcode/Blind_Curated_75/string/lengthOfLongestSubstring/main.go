package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	idxLeft := 0
	maxLength := 0
	currLength := 0
	char2Idx := map[rune]int{}

	for idxRight, currByte := range s {
		lastIdx, found := char2Idx[currByte]
		// if found repeat and lastIdx >= idxLeft(must in the range of idxLeft and idxRight)
		if found && lastIdx >= idxLeft {
			idxLeft = lastIdx + 1
		} else {
			currLength = idxRight - idxLeft + 1
			if maxLength < currLength {
				maxLength = currLength
			}
		}
		char2Idx[currByte] = idxRight
	}
	return maxLength
}

func main() {
	s := "abcabcbb"
	k := lengthOfLongestSubstring(s)
	fmt.Printf("k:%+v\n", k)
}
