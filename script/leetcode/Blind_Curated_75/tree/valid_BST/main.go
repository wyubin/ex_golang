package main

import "fmt"

func validNodeRegion(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return validNodeRegion(root.Left, min, root.Val) && validNodeRegion(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return validNodeRegion(root, 0, 10001)
}

func main() {
	s := []int{2, 1, 4, -1, -1, 3, 6}
	tree := createTree(s)
	k := isValidBST(tree)
	fmt.Printf("k:%+v\n", k)
}
