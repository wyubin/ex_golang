package main

import "fmt"

const nodesSize = 100

func cloneGraph(node *Node) *Node {
	node2clone := map[*Node]*Node{}
	ch := make(chan *Node, nodesSize)
	defer close(ch)

	// init root node clone
	rootNode := &Node{Val: node.Val}
	node2clone[node] = rootNode
	ch <- node

	for len(ch) > 0 {
		currNode := <-ch
		for _, neighbor := range currNode.Neighbors {
			if _, found := node2clone[neighbor]; !found {
				clone := &Node{Val: neighbor.Val}
				node2clone[neighbor] = clone
				ch <- neighbor
			}
			node2clone[currNode].Neighbors = append(node2clone[currNode].Neighbors, node2clone[neighbor])
		}
	}
	return rootNode
}

func main() {
	s := [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}
	k := cloneGraph(makeGraph(s))
	fmt.Printf("k:%+v\n", k)
}
