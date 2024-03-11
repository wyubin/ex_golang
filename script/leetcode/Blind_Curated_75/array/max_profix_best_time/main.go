package main

import "fmt"

func solution(prices []int) int {
	maxProfix, idxBuy, idxSell := 0, 0, 1
	var currProfix int
	for idxSell < len(prices) {
		if prices[idxSell] > prices[idxBuy] {
			currProfix = prices[idxSell] - prices[idxBuy]
			if currProfix > maxProfix {
				maxProfix = currProfix
			}
		} else {
			// if <= then buy
			idxBuy = idxSell
		}
		idxSell++
	}
	return maxProfix
}

func main() {
	s := []int{7, 1, 5, 3, 6, 4}
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
