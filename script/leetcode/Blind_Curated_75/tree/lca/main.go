package main

import "fmt"

func lcaRetrive(rootNode *TreeNode, p, q int) *TreeNode {
	if rootNode.Val > q {
		return lcaRetrive(rootNode.Left, p, q)
	} else if rootNode.Val < p {
		return lcaRetrive(rootNode.Right, p, q)
	}
	return rootNode
}

func solution(inTree *TreeNode, p, q int) int {
	if q < p {
		p, q = q, p
	}
	lcaNode := lcaRetrive(inTree, p, q)
	return lcaNode.Val
}

func main() {
	s := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	k := solution(createTree(s), 2, 8)
	fmt.Printf("k: %+v\n", k)
}
