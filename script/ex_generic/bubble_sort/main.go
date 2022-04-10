package main

import (
	"fmt"

	constraints "ex_golang/lib/contraints"
)

func bubbleSortGeneric[T constraints.Ordered](x []T) {
	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func main() {
	val := []int{1, 3, 2}
	bubbleSortGeneric(val)
	fmt.Printf("val:%+v", val)
}
