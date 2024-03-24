package main

import (
	"fmt"
	"slices"
)

func solution(srcStr string, words []string) bool {
	// init breakpoint
	dp := make([]bool, len(srcStr)+1)
	dp[0] = true
	maxWordLen := 0
	for _, word := range words {
		currLen := len(word)
		if maxWordLen < currLen {
			maxWordLen = currLen
		}
	}

	for idxEnd := 1; idxEnd <= len(srcStr); idxEnd++ {
		maxStart := slices.Max([]int{0, idxEnd - maxWordLen})
		// check dp[idxEnd] can be true
		for idxStart := idxEnd - 1; idxStart >= maxStart; idxStart-- {
			if dp[idxStart] && slices.Contains(words, srcStr[idxStart:idxEnd]) {
				dp[idxEnd] = true
				break
			}
		}
	}
	return dp[len(srcStr)]
}

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	k := solution(s, wordDict)
	fmt.Printf("k:%+v\n", k)
}
