package main

import "fmt"

func posExtend(grid [][]string, idxRow, idxCol int, visitedPos map[string]struct{}) bool {
	if idxRow < 0 || idxRow >= len(grid) || idxCol < 0 || idxCol >= len(grid[0]) {
		return false
	}
	if grid[idxRow][idxCol] == "0" {
		return false
	}
	strPos := fmt.Sprintf("%d_%d", idxRow, idxCol)
	fmt.Printf("try %s\n", strPos)
	if _, found := visitedPos[strPos]; found {
		return false
	}
	fmt.Printf("visit island with %s\n", strPos)
	visitedPos[strPos] = struct{}{}
	// entend
	posExtend(grid, idxRow-1, idxCol, visitedPos)
	posExtend(grid, idxRow+1, idxCol, visitedPos)
	posExtend(grid, idxRow, idxCol-1, visitedPos)
	posExtend(grid, idxRow, idxCol+1, visitedPos)
	return true
}

func solution(grid [][]string) int {
	countIslands := 0
	visitedPos := map[string]struct{}{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if posExtend(grid, i, j, visitedPos) {
				countIslands++
			}
		}
	}
	return countIslands
}

func main() {
	s := [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "1", "0", "0"},
		{"0", "0", "0", "1", "1"},
	}
	k := solution(s)
	fmt.Printf("k: %d\n", k)
}
