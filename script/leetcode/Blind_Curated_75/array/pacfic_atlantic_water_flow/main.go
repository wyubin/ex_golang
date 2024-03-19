package main

import "fmt"

var (
	pacDP = [][]bool{}
	atlDP = [][]bool{}
)

func dfs(matH [][]int, visited [][]bool, x, y int, res *[][]int) {
	if visited[x][y] {
		return
	}
	fmt.Printf("visited: [%d,%d]\n", x, y)
	visited[x][y] = true
	if pacDP[x][y] && atlDP[x][y] {
		*res = append(*res, []int{x, y})
	}
	// retrieve
	currH := matH[x][y]
	up, down, left, right := y-1, y+1, x-1, x+1
	numRow, numCol := len(matH), len(matH[0])
	if up >= 0 && currH <= matH[x][up] {
		dfs(matH, visited, x, up, res)
	}
	if down < numRow && currH <= matH[x][down] {
		dfs(matH, visited, x, down, res)
	}
	if left >= 0 && currH <= matH[left][y] {
		dfs(matH, visited, left, y, res)
	}
	if right < numCol && currH <= matH[right][y] {
		dfs(matH, visited, right, y, res)
	}
}

func solution(matH [][]int) [][]int {
	// init DPs
	numCol := len(matH)
	numRow := len(matH[0])
	for i := 0; i < numCol; i++ {
		pacDP = append(pacDP, make([]bool, numRow))
		atlDP = append(atlDP, make([]bool, numRow))
	}
	res := [][]int{}
	// retrieve from ocean
	for i := 0; i < numCol; i++ {
		dfs(matH, pacDP, i, 0, &res)
		dfs(matH, atlDP, i, numRow-1, &res)
	}
	for j := 0; j < numRow; j++ {
		dfs(matH, pacDP, 0, j, &res)
		dfs(matH, atlDP, numCol-1, j, &res)
	}
	return res
}

func main() {
	matH := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	k := solution(matH)
	fmt.Printf("k: %+v\n", k)
}
