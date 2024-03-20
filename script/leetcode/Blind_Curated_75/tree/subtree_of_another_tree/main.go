package main

import "fmt"

func findSubNode(rootTree, subTree *TreeNode) *TreeNode {
	if rootTree == nil || rootTree.Val == subTree.Val {
		return rootTree
	}
	resLeft := findSubNode(rootTree.Left, subTree)
	if resLeft == nil {
		return findSubNode(rootTree.Right, subTree)
	}
	return resLeft
}

func isSameTree(rootTree, subTree *TreeNode) bool {
	if rootTree == nil || subTree == nil {
		return rootTree == subTree
	}
	if rootTree.Val != subTree.Val {
		return false
	}
	return isSameTree(rootTree.Left, subTree.Left) && isSameTree(rootTree.Right, subTree.Right)
}

func solution(rootTree, subTree *TreeNode) bool {
	subNode := findSubNode(rootTree, subTree)
	if subNode == nil {
		return false
	}
	return isSameTree(subNode, subTree)
}

func main() {
	root := []int{3, 4, 5, 1, 2}
	sub := []int{4, 1, 2}
	k := solution(createTree(root), createTree(sub))
	fmt.Printf("k: %+v\n", k)
}
