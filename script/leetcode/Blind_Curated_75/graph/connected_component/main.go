package main

import "fmt"

func triverseNodes(node2childs map[int][]int, visited map[int]bool, nodeStart int) {
	if visited[nodeStart] {
		return
	}
	visited[nodeStart] = true
	for _, child := range node2childs[nodeStart] {
		triverseNodes(node2childs, visited, child)
	}
}

func solution(countNodes int, edges [][]int) int {
	visited := make(map[int]bool, countNodes)
	node2childs := map[int][]int{}
	for _, edge := range edges {
		// bi-direction
		node2childs[edge[0]] = append(node2childs[edge[0]], edge[1])
		node2childs[edge[1]] = append(node2childs[edge[1]], edge[0])
	}
	// start dfs
	count := 0
	for i := 0; i < countNodes; i++ {
		if !visited[i] {
			count++
			triverseNodes(node2childs, visited, i)
		}
	}
	return count
}

func main() {
	n := 5
	edges := [][]int{
		{0, 1}, {1, 2}, {3, 4},
	}
	k := solution(n, edges)
	fmt.Printf("k: %+v\n", k)
}
