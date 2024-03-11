package main

import (
	"fmt"
	"slices"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return slices.Max([]int{maxDepth(root.Left), maxDepth(root.Right)}) + 1
}

func main() {
	s := []int{3, 9, 20, -1, -1, 15, 7}
	tree := createTree(s)
	k := maxDepth(tree)
	fmt.Printf("k:%+v\n", k)
}
