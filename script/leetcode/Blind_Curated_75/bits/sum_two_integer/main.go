package main

import "fmt"

func solution(a, b int) int {
	fmt.Printf("start a[%b], b[%b]\n", a, b)
	for b != 0 {
		carry := a & b
		fmt.Printf("make carry: %b\n", carry)
		a = a ^ b
		b = carry << 1
		fmt.Printf("processing a[%b], b[%b]\n", a, b)
	}
	return a
}

func main() {
	a, b := 3, 5
	k := solution(a, b)
	fmt.Printf("k: %+v\n", k)
}
