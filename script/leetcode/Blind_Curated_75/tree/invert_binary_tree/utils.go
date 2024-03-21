package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (s *TreeNode) String() string {
	if s == nil {
		return "[]"
	}
	res := []string{}
	queue := []*TreeNode{s}
	var currNode *TreeNode
	for len(queue) > 0 {
		currNode, queue = queue[0], queue[1:]
		if currNode == nil {
			res = append(res, "-1")
			continue
		}
		res = append(res, strconv.Itoa(currNode.Val))
		queue = append(queue, currNode.Left, currNode.Right)
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] != "-1" {
			res = res[:(i + 1)]
			break
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ","))
}

func createTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[0]
	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] != -1 {
			tree.Left = new(TreeNode)
			tree.Left.Val = nums[i]
			ch <- tree.Left
		}
		i++
		if i < len(nums) || nums[i] != -1 {
			tree.Right = new(TreeNode)
			tree.Right.Val = nums[i]
			ch <- tree.Right
		}
	}
	close(ch)
	return root
}
