package main

import (
	"fmt"
	"strconv"
)

func solution(bs string) string {
	bsInt64, err := strconv.ParseUint(bs, 2, 32)
	if err != nil {
		return err.Error()
	}
	bsInt32 := uint32(bsInt64)
	res := uint32(0)
	for i := 0; i < 32 && bsInt32 != 0; i++ {
		currBit := bsInt32 & 1
		revBit := currBit << (31 - i)
		res = res | revBit
		// fmt.Printf("shift %d %dth time\n", bsInt32, i+1)
		bsInt32 = bsInt32 >> 1
	}
	return fmt.Sprintf("%032b", res)
}

func main() {
	s := "00000010100101000001111010011100"
	k := solution(s)
	fmt.Printf("k: %s\n", k)
}
