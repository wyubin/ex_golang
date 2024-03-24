package main

import "fmt"

var (
	pacLand = [][]bool{}
	atlLand = [][]bool{}
)

func waterTrace(visited [][]bool, matLand [][]int, idxRow, idxCol int, posSame *[][]int) {
	nRow, nCol := len(matLand), len(matLand[0])
	if visited[idxRow][idxCol] {
		return
	}
	visited[idxRow][idxCol] = true
	// check pos same
	if pacLand[idxRow][idxCol] && atlLand[idxRow][idxCol] {
		*posSame = append(*posSame, []int{idxRow, idxCol})
	}
	// trace next
	currH := matLand[idxRow][idxCol]
	up, down, left, right := idxRow-1, idxRow+1, idxCol-1, idxCol+1
	if up >= 0 && currH <= matLand[up][idxCol] {
		waterTrace(visited, matLand, up, idxCol, posSame)
	}
	if down < nRow && currH <= matLand[down][idxCol] {
		waterTrace(visited, matLand, down, idxCol, posSame)
	}
	if left >= 0 && currH <= matLand[idxRow][left] {
		waterTrace(visited, matLand, idxRow, left, posSame)
	}
	if right < nCol && currH <= matLand[idxRow][right] {
		waterTrace(visited, matLand, idxRow, right, posSame)
	}
}

func solution(matLand [][]int) [][]int {
	// init pacLand, atlLand
	nRow, nCol := len(matLand), len(matLand[0])
	for i := 0; i < nRow; i++ {
		pacLand = append(pacLand, make([]bool, nCol))
		atlLand = append(atlLand, make([]bool, nCol))
	}
	// use waterTrace
	posSame := [][]int{}
	for idxRow := 0; idxRow < nRow; idxRow++ {
		waterTrace(pacLand, matLand, idxRow, 0, &posSame)
		waterTrace(atlLand, matLand, idxRow, nCol-1, &posSame)
	}
	for idxCol := 0; idxCol < nCol; idxCol++ {
		waterTrace(pacLand, matLand, 0, idxCol, &posSame)
		waterTrace(atlLand, matLand, nRow-1, idxCol, &posSame)
	}
	return posSame
}

func main() {
	matH := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	k := solution(matH)
	fmt.Printf("k: %+v\n", k)
}
