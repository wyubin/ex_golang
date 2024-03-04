package main

import "fmt"

func solution(s string) string {
	// make 2d slice with false
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	maxLength := 1
	maxSubstr := s[0:1]
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		for j := 0; j < i; j++ {
			// check palindrome based on dp
			if s[j] == s[i] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if i-j+1 > maxLength {
					maxLength = i - j + 1
					maxSubstr = s[j : i+1]
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
