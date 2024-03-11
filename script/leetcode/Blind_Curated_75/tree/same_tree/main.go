package main

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	s := []int{1, 2, 3, 4}
	t := []int{1, 2, 3, 4}
	k := isSameTree(createTree(s), createTree(t))
	fmt.Printf("k:%+v\n", k)
}
