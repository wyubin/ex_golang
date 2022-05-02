package easyother

func isSymmetric(root *TreeNode) bool {
	return isSymmetricPair(root.Left, root.Right)
}

func isSymmetricPair(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	return isSymmetricPair(left.Left, right.Right) && isSymmetricPair(right.Left, left.Right)
}
