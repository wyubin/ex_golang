package main

import "fmt"

func dumpLevel(root *TreeNode, res [][]int, lv int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == lv {
		// init lv slice
		res = append(res, []int{})
	}
	res[lv] = append(res[lv], root.Val)
	res = dumpLevel(root.Left, res, lv+1)
	res = dumpLevel(root.Right, res, lv+1)
	return res
}

func solution(root *TreeNode) [][]int {
	return dumpLevel(root, [][]int{}, 0)
}

func main() {
	s := []int{3, 9, 20, -1, -1, 15, 7}
	tree := createTree(s)
	k := solution(tree)
	fmt.Printf("k:%+v\n", k)
}
