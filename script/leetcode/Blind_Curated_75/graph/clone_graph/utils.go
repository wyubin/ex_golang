package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func makeGraph(pairs [][]int) *Node {
	if len(pairs) == 0 {
		return nil
	}
	nodes := make(map[int]*Node)
	for i := 1; i <= len(pairs); i++ {
		nodes[i] = &Node{Val: i}
	}
	for idx, pair := range pairs {
		neighbors := []*Node{}
		for _, val := range pair {
			neighbors = append(neighbors, nodes[val])
		}
		nodes[idx+1].Neighbors = neighbors
	}
	return nodes[1]
}
