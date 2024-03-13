package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) String() string {
	strContent := ""
	visited := map[*ListNode]bool{}
	for node := s; node != nil && !visited[node]; node = node.Next {
		if strContent != "" {
			strContent += ","
		}
		strContent += fmt.Sprintf("%d", node.Val)
		visited[node] = true
	}
	return fmt.Sprintf("[%s]", strContent)
}

func slice2ListNode(nums []int, posTailNext int) *ListNode {
	head := &ListNode{}
	curr := head
	var nodeTailNext *ListNode
	for idx, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
		if idx == posTailNext {
			nodeTailNext = curr
		}
	}
	if nodeTailNext != nil {
		curr.Next = nodeTailNext
	}
	return head.Next
}
