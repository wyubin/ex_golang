package main

import "fmt"

func solution(s, t string) string {
	byte2count := map[byte]int{}
	count := 0
	for _, char := range t {
		byte2count[byte(char)]++
		count++
	}
	start, end := 0, 0
	minLen, start_idx := len(s)+1, 0
	for end < len(s) {
		_, found := byte2count[s[end]]
		if found {
			byte2count[s[end]]--
			count--
		}
		end++

		// move start
		for count == 0 {
			if end-start < minLen {
				minLen = end - start
				start_idx = start
			}
			_, found := byte2count[s[start]]
			if found {
				byte2count[s[start]]++
				count++
			}
			start++
		}
	}
	if minLen < len(s)+1 {
		return s[start_idx : start_idx+minLen]
	}
	return ""
}

func main() {
	s := "a"
	t := "aa"
	k := solution(s, t)
	fmt.Printf("k:%+v\n", k)
}
