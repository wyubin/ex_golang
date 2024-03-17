package main

import "fmt"

func invertTree(inTree *TreeNode) {
	if inTree == nil {
		return
	}
	invertTree(inTree.Left)
	invertTree(inTree.Right)
	inTree.Left, inTree.Right = inTree.Right, inTree.Left
}

func main() {
	s := []int{4, 2, 7, 1, 3, 6, 9}
	treeS := createTree(s)
	fmt.Printf("treeS: %s\n", treeS)
	invertTree(treeS)
	fmt.Printf("invert treeS: %s\n", treeS)
}
