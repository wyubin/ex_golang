package main

import "fmt"

func buildSubtree(preorder []int, inorder2idx map[int]int, rootIdx, idxLast, inorderStart, inorderEnd int) *TreeNode {
	if inorderStart > inorderEnd || rootIdx > idxLast {
		return nil
	}
	// create Node
	root := new(TreeNode)
	root.Val = preorder[rootIdx]
	// left and right build
	midIdx := inorder2idx[root.Val]
	lenLeft := midIdx - inorderStart
	fmt.Printf("currIdx:%d, inorderStart:%d, inorderEnd:%d, midIdx: %d, lenLeft: %d\n", rootIdx, inorderStart, inorderEnd, midIdx, lenLeft)
	root.Left = buildSubtree(preorder, inorder2idx, rootIdx+1, rootIdx+lenLeft, inorderStart, midIdx-1)
	root.Right = buildSubtree(preorder, inorder2idx, rootIdx+lenLeft+1, idxLast, midIdx+1, inorderEnd)
	return root
}

func buildTree(preorder, inorder []int) *TreeNode {
	inorder2idx := make(map[int]int)
	for idx, val := range inorder {
		inorder2idx[val] = idx
	}
	currIdx := 0
	return buildSubtree(preorder, inorder2idx, currIdx, len(preorder)-1, 0, len(inorder)-1)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	k := buildTree(preorder, inorder)

	fmt.Printf("k:%s\n", k)
}
