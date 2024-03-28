package main

import "fmt"

func dumpPaths(dp [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	paths := dp[i][j]
	if paths != 0 {
		return paths
	}
	dp[i][j] = dumpPaths(dp, i-1, j) + dumpPaths(dp, i, j-1)
	return dp[i][j]
}

func uniquePaths(m int, n int) int {
	// initial dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// dumpPaths
	res := dumpPaths(dp, m-1, n-1)
	fmt.Printf("dp:%+v\n", dp)
	return res
}

func main() {
	k := uniquePaths(3, 7)
	fmt.Printf("k:%+v\n", k)
}
