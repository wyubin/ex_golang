package main

import (
	"fmt"
	"slices"
	"strings"
)

func dfs(board [][]string, chars []string, i, j, k int, paths *[]string) bool {
	if k == len(chars) {
		return true
	}
	currPath := fmt.Sprintf("%d_%d", i, j)
	if slices.Index(*paths, currPath) != -1 {
		return false
	}
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[i]) ||
		board[i][j] != chars[k] {
		return false
	}
	*paths = append(*paths, currPath)
	if dfs(board, chars, i+1, j, k+1, paths) ||
		dfs(board, chars, i-1, j, k+1, paths) ||
		dfs(board, chars, i, j+1, k+1, paths) ||
		dfs(board, chars, i, j-1, k+1, paths) {
		return true
	}
	*paths = (*paths)[:len(*paths)-1]
	return false
}

func solution(board [][]string, word string) bool {
	chars := strings.Split(word, "")
	paths := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == chars[0] {
				if dfs(board, chars, i, j, 0, &paths) {
					fmt.Printf("paths: %+v\n", paths)
					return true
				}
			}
		}
	}
	return false
}

func main() {
	board := [][]string{
		{"A", "B", "C", "E"},
		{"S", "F", "C", "S"},
		{"A", "D", "E", "E"},
	}
	word := "ABCCED"
	fmt.Println(solution(board, word))
	// ext, word search 2
	board = [][]string{
		{"o", "a", "a", "n"},
		{"e", "t", "a", "e"},
		{"i", "h", "k", "r"},
	}
	for _, word := range []string{"oath", "pea", "eat", "rain"} {
		k := solution(board, word)
		fmt.Printf("search[%s]: %+v\n", word, k)
	}
}
