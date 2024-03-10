package main

import (
	"fmt"
	"strconv"
)

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
	for i := 2; i <= len(numstr); i++ {
		code1 := numstr[i-1 : i]
		code2 := numstr[i-2 : i]
		if _, ok := code[code1]; ok {
			dp[i] += dp[i-1]
		}
		if _, ok := code[code2]; ok {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(numstr)]
}
func main() {
	s := "226619"
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
