package easyother

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val  *TreeNode
	Next *Node
}

type BinaryStack struct {
	top *Node
}

func (s *BinaryStack) Push(a *TreeNode) {
	next := s.top
	s.top = &Node{
		Val:  a,
		Next: next,
	}
}

func (s *BinaryStack) Pop() *TreeNode {
	if s.top == nil {
		return nil
	}
	node := s.top
	s.top = node.Next
	return node.Val
}

func inorderTraversal(root *TreeNode) []int {
	node := root

	arr := []int{}
	stack := &BinaryStack{}
	popped := false

	for node != nil {
		if node.Left != nil && !popped {
			stack.Push(node)
			node = node.Left
			popped = false
		} else {
			arr = append(arr, node.Val)
			if node.Right != nil {
				node = node.Right
				popped = false
			} else {
				node = stack.Pop()
				popped = true
			}
		}
	}
	return arr
}
