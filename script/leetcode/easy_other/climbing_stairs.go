package easyother

// recursive but slow
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	} else if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	dp := []int{0, 1, 2}
	if n < 3 {
		return dp[n]
	}
	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
