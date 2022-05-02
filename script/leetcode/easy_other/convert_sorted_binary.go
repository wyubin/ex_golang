package easyother

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 計算 mid
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[0:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
