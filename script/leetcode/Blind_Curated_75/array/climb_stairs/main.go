package main

import "fmt"

func dumpSteps(distance int, dp map[int]int) int {
	if distance == 0 || distance == 1 {
		return 1
	}
	stepsTmp, found := dp[distance]
	if found {
		return stepsTmp
	}
	dp[distance] = dumpSteps(distance-1, dp) + dumpSteps(distance-2, dp)
	return dp[distance]
}

func climbStairs(distance int) int {
	dp := map[int]int{}
	return dumpSteps(distance, dp)
}

func main() {
	num := 4
	k := climbStairs(num)
	fmt.Printf("k:%+v\n", k)
}
