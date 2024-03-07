package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	idxTop, idxBottom, idxLeft, idxRight := 0, len(matrix)-1, 0, len(matrix[0])-1
	for idxTop <= idxBottom && idxLeft <= idxRight {
		// top
		for i := idxLeft; i <= idxRight; i++ {
			res = append(res, matrix[idxTop][i])
		}
		idxTop++
		// right
		for i := idxTop; i <= idxBottom; i++ {
			res = append(res, matrix[i][idxRight])
		}
		idxRight--
		// bottom
		if idxTop <= idxBottom {
			for i := idxRight; i >= idxLeft; i-- {
				res = append(res, matrix[idxBottom][i])
			}
		}
		idxBottom--
		// left
		if idxLeft <= idxRight {
			for i := idxBottom; i >= idxTop; i-- {
				res = append(res, matrix[i][idxLeft])
			}
		}
		idxLeft++
	}
	return res
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	k := spiralOrder(mat)
	fmt.Printf("k:%+v\n", k)
}
