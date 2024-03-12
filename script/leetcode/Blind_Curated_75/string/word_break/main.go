package main

import (
	"fmt"
	"slices"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	maxWord := 0
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
		if len(word) > maxWord {
			maxWord = len(word)
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= slices.Max([]int{i - 1 - maxWord, 0}); j-- {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	k := wordBreak(s, wordDict)
	fmt.Printf("k:%+v\n", k)
}
