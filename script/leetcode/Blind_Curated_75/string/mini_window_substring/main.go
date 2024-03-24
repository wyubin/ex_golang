package main

import "fmt"

func solution(srcStr, trgStr string) string {
	// init ref data
	countNeed := 0
	byte2count := map[byte]int{}
	for _, char := range trgStr {
		countNeed++
		byte2count[byte(char)]++
	}
	// init pointer
	minLen, minStart := len(srcStr)+1, 0
	idxStart, idxEnd := 0, 0
	for idxEnd < len(srcStr) {
		byteEnd := srcStr[idxEnd]
		_, found := byte2count[byteEnd]
		if found {
			byte2count[byteEnd]--
			countNeed--
		}
		idxEnd++
		// if countNeed == 0
		for countNeed == 0 {
			currLen := idxEnd - idxStart
			if currLen < minLen {
				minLen = currLen
				minStart = idxStart
			}
			byteStart := srcStr[idxStart]
			_, found := byte2count[byteStart]
			if found {
				byte2count[byteStart]++
				countNeed++
			}
			idxStart++
		}
	}
	if minLen < len(srcStr)+1 {
		return srcStr[minStart : minStart+minLen]
	}
	return ""
}

func main() {
	s := "a"
	t := "aa"
	k := solution(s, t)
	fmt.Printf("k:%+v\n", k)
}
