package main

import (
	"fmt"
)

func rotate(mat [][]int) {
	// transpose
	for i := 0; i < len(mat); i++ {
		for j := i + 1; j < len(mat[0]); j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	// col swap
	for j := 0; j < len(mat); j++ {
		idxStart, idxEnd := 0, len(mat[0])-1
		for idxStart < idxEnd {
			mat[j][idxStart], mat[j][idxEnd] = mat[j][idxEnd], mat[j][idxStart]
			idxStart++
			idxEnd--
		}
	}
}

func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(nums)
	fmt.Printf("nums:%+v\n", nums)
}
