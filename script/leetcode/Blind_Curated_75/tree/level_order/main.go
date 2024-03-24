package main

import "fmt"

func dumpLevel(root *TreeNode, res *[][]int, lv int) {
	if root == nil {
		return
	}
	if len(*res) == lv {
		// init lv slice
		*res = append(*res, []int{})
	}
	(*res)[lv] = append((*res)[lv], root.Val)
	dumpLevel(root.Left, res, lv+1)
	dumpLevel(root.Right, res, lv+1)
}

func solution(root *TreeNode) [][]int {
	res := [][]int{}
	dumpLevel(root, &res, 0)
	return res
}

func main() {
	s := []int{3, 9, 20, -1, -1, 15, 7}
	tree := createTree(s)
	k := solution(tree)
	fmt.Printf("k:%+v\n", k)
}
