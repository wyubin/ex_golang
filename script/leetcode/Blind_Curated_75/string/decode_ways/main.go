package main

import (
	"fmt"
	"strconv"
)

func dumpDecode(encStr string, idx int, code map[string]struct{}, dp map[int]int) int {
	if idx < 0 {
		return 0
	}
	numDecode, found := dp[idx]
	if found {
		return numDecode
	}
	code1 := encStr[idx-1 : idx]
	code2 := encStr[idx-2 : idx]
	dp[idx] = 0
	if _, ok := code[code1]; ok {
		dp[idx] += dumpDecode(encStr, idx-1, code, dp)
	}
	if _, ok := code[code2]; ok {
		dp[idx] += dumpDecode(encStr, idx-2, code, dp)
	}
	return dp[idx]
}

func solution(numstr string) int {
	dp := map[int]int{}
	code := map[string]struct{}{}
	for i := 1; i < 27; i++ {
		code[strconv.Itoa(i)] = struct{}{}
	}
	if string(numstr[0]) == "0" {
		return 0
	}
	dp[0] = 1
	dp[1] = 1
	return dumpDecode(numstr, len(numstr), code, dp)
}
func main() {
	s := "226"
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
