package main

import "fmt"

func solution(inStr string) int {
	dp := [][]bool{}
	lenStr := len(inStr)
	for i := 0; i < lenStr; i++ {
		dp = append(dp, make([]bool, lenStr))
	}
	res := []string{}
	for i := 0; i < lenStr; i++ {
		dp[i][i] = true // init single palindromic
		res = append(res, inStr[i:i+1])
		for j := 0; j < i; j++ {
			if inStr[j] == inStr[i] && (i-j == 1 || dp[j+1][i-1]) {
				dp[j][i] = true
				res = append(res, inStr[j:(i+1)])
			}
		}
	}
	fmt.Printf("res: %+v\n", res)
	return len(res)
}

func main() {
	s := "aaa"
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
