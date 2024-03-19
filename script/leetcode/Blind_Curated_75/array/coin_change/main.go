package main

import (
	"fmt"
	"math"
	"slices"
)

const MAXINT = math.MaxUint32

func minCoinCount(dp map[string]int, coins []int, idxCoin int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return MAXINT
	}
	if idxCoin == 0 {
		if amount%coins[idxCoin] == 0 {
			return amount / coins[idxCoin]
		}
		return MAXINT
	}
	keyCoinAmount := fmt.Sprintf("%d_%d", idxCoin, amount)
	if dp[keyCoinAmount] != 0 {
		return dp[keyCoinAmount]
	}
	useCurrCoin := 1 + minCoinCount(dp, coins, idxCoin, amount-coins[idxCoin])
	nextCoin := minCoinCount(dp, coins, idxCoin-1, amount)
	dp[keyCoinAmount] = slices.Min([]int{useCurrCoin, nextCoin})
	return dp[keyCoinAmount]
}

func solution(coins []int, amount int) int {
	dp := map[string]int{}
	res := minCoinCount(dp, coins, len(coins)-1, amount)
	if res == MAXINT {
		return -1
	}
	return res
}

func main() {
	coins := []int{2}
	amount := 3
	k := solution(coins, amount)
	fmt.Printf("k: %+v\n", k)
}
