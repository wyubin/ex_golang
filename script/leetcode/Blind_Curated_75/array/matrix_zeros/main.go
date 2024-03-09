package main

import "fmt"

func solution(mat [][]int) {
	zCol, zRow := map[int]bool{}, map[int]bool{}
	for i, row := range mat {
		for j, val := range row {
			if val == 0 {
				zCol[j] = true
				zRow[i] = true
			}
		}
	}
	// set to zero
	for i, row := range mat {
		for j := range row {
			if zCol[j] || zRow[i] {
				mat[i][j] = 0
			}
		}
	}
}

func main() {
	mat := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	solution(mat)
	fmt.Printf("mat:%+v\n", mat)
}
