package main

import "fmt"

func solution(lastNum int) []int {
	res := []int{}
	for i := 0; i < lastNum+1; i++ {
		tmpNum := i
		countBit := 0
		for tmpNum != 0 {
			if tmpNum&1 == 1 {
				countBit++
			}
			tmpNum >>= 1
		}
		res = append(res, countBit)
	}
	return res
}

func main() {
	n := 5
	k := solution(n)
	fmt.Printf("k: %+v\n", k)
}
