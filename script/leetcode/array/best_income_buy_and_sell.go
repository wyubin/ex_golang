package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	lenPrice := len(prices)
	valBuy := -1
	for ind, val := range prices {
		if ind+1 == lenPrice || val > prices[ind+1] {
			if valBuy > -1 {
				profit += val - valBuy
				fmt.Printf("sale out(%d - %d):%d day\n", valBuy, val, ind)
				valBuy = -1
			}
		} else if valBuy == -1 {
			valBuy = val
			fmt.Printf("buy in:%d day\n", ind)
		}
	}
	return profit
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	k := maxProfit(nums)
	//fmt.Printf("nums:%+v\n", nums)
	fmt.Printf("k:%+v\n", k)
}
