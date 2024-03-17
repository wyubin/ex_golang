package main

import "fmt"

func inOrderRetrive(inTree *TreeNode, inOrderNodes *[]int) {
	if inTree == nil {
		return
	}
	// append left, then root, then, right
	inOrderRetrive(inTree.Left, inOrderNodes)
	*inOrderNodes = append(*inOrderNodes, inTree.Val)
	fmt.Printf("Append node:%d\n", inTree.Val)
	inOrderRetrive(inTree.Right, inOrderNodes)
}

func solution(inTree *TreeNode, idxBased1 int) int {
	inOrderNodes := []int{}
	inOrderRetrive(inTree, &inOrderNodes)
	return inOrderNodes[idxBased1-1]
}

func main() {
	s := []int{5, 3, 6, 2, 4, -1, -1, 1}
	i := 3
	k := solution(createTree(s), i)
	fmt.Printf("k: %+v\n", k)
}
