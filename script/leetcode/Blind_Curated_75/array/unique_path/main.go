package main

import "fmt"

func dumpPaths(m int, n int, i int, j int, dp [][]int) int {
	if i >= m || j >= n {
		return 0
	}
	if i == m-1 && j == n-1 {
		// to the end
		return 1
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	dp[i][j] = dumpPaths(m, n, i+1, j, dp) + dumpPaths(m, n, i, j+1, dp)
	return dp[i][j]
}

func uniquePaths(m int, n int) int {
	// initial dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// dumpPaths
	res := dumpPaths(m, n, 0, 0, dp)
	fmt.Printf("dp:%+v\n", dp)
	return res
}

func main() {
	k := uniquePaths(3, 7)
	fmt.Printf("k:%+v\n", k)
}
