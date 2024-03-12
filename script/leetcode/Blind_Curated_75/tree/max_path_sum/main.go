package main

import (
	"fmt"
	"slices"
)

// return single branch max and dump localmax
func maxGainDump(node *TreeNode, maxVal *int) int {
	if node == nil {
		return 0
	}

	leftGain := slices.Max([]int{maxGainDump(node.Left, maxVal), 0})
	rightGain := slices.Max([]int{maxGainDump(node.Right, maxVal), 0})
	branchMax := node.Val + slices.Max([]int{leftGain, rightGain})

	// update local max
	currMaxSum := node.Val + leftGain + rightGain
	*maxVal = slices.Max([]int{*maxVal, currMaxSum})
	return branchMax
}

func solution(root *TreeNode) int {
	maxVal := 0
	maxGainDump(root, &maxVal)
	return maxVal
}

func main() {
	s := []int{-10, 9, 20, -1, -1, 15, 7}
	k := solution(createTree(s))
	fmt.Printf("k:%+v\n", k)
}
