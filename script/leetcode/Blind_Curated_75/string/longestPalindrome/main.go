package main

import "fmt"

func solution(inStr string) string {
	pos2Pal := [][]bool{}
	lenStr := len(inStr)
	for i := 0; i < lenStr; i++ {
		pos2Pal = append(pos2Pal, make([]bool, lenStr))
	}
	maxLength := 1
	currLength := 0
	maxSubstr := inStr[0:1]
	for idxEnd := 0; idxEnd < lenStr; idxEnd++ {
		pos2Pal[idxEnd][idxEnd] = true
		for idxStart := 0; idxStart < idxEnd; idxStart++ {
			currLength = idxEnd - idxStart + 1
			if inStr[idxStart] == inStr[idxEnd] && (currLength == 2 || pos2Pal[idxStart+1][idxEnd-1]) {
				pos2Pal[idxStart][idxEnd] = true
				if currLength > maxLength {
					maxLength = currLength
					maxSubstr = inStr[idxStart : idxEnd+1]
				}
			}
		}
	}
	return maxSubstr
}

func main() {
	s := "babad"
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
