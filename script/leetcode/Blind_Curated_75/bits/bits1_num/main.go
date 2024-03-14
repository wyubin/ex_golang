package main

import (
	"fmt"
	"strconv"
)

func solution(bs string) int {
	bsInt, err := strconv.ParseUint(bs, 2, 32)
	if err != nil {
		panic(err)
	}
	oneLen := 0
	for bsInt != 0 {
		if bsInt&1 == 1 {
			oneLen++
		}
		bsInt >>= 1
	}
	return oneLen
}

func main() {
	s := "00000000000000000000000010000000"
	k := solution(s)
	fmt.Printf("k: %d\n", k)
}
