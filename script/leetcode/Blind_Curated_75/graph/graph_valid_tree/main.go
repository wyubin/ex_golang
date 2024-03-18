package main

import "fmt"

func checkCircle(graph map[int][]int, visited map[int]bool, currNode, preNode int) bool {
	if visited[currNode] {
		return false
	}
	visited[currNode] = true
	for _, childNode := range graph[currNode] {
		if childNode == preNode {
			continue
		}
		if !checkCircle(graph, visited, childNode, currNode) {
			return false
		}
	}
	return true
}

func solution(nodeSize int, edges [][]int) bool {
	node2child := map[int][]int{}
	visited := map[int]bool{}
	for _, edge := range edges {
		// bi-direction
		node2child[edge[0]] = append(node2child[edge[0]], edge[1])
		node2child[edge[1]] = append(node2child[edge[1]], edge[0])
	}
	// check circle
	if !checkCircle(node2child, visited, 0, -1) {
		return false
	}
	// check all visited
	for i := 0; i < nodeSize; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}

func main() {
	s := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	k := solution(len(s), s)
	fmt.Printf("k: %+v\n", k)
	s = [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	k = solution(len(s), s)
	fmt.Printf("k: %+v\n", k)
}
