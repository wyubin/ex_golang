package main

import "fmt"

func reverseString(s []byte) {
	lenS := len(s) - 1
	for ind := 0; ind < (lenS+1)/2; ind++ {
		oppoInd := lenS - ind
		s[ind], s[oppoInd] = s[oppoInd], s[ind]
	}
}

func main() {
	nums := []byte("string")
	reverseString(nums)
	fmt.Printf("nums:%s\n", string(nums))
}
