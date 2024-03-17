package main

import "fmt"

func solution(a, b string) bool {
	byte2len := map[byte]int{}
	for i := 0; i < len(a); i++ {
		byte2len[a[i]]++
		byte2len[b[i]]--
	}
	for _, count := range byte2len {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "rat"
	t := "car"
	k := solution(s, t)
	fmt.Printf("k: %+v\n", k)
}
