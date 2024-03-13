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

func slice2ListNode(nums []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return head.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	nodeLen := 0
	node := head
	for node != nil {
		nodeLen++
		node = node.Next
	}
	currIdx := 0
	idxJump := nodeLen - n - 1
	if idxJump < 0 {
		return head.Next
	}
	node = head
	for node != nil {
		if currIdx == idxJump {
			node.Next = node.Next.Next
			break
		}
		currIdx++
		node = node.Next
	}
	return dummy.Next
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	n := 2
	head := slice2ListNode(nums)
	fmt.Printf("head:%s\n", head)
	k := removeNthFromEnd(head, n)
	fmt.Printf("k:%s\n", k)
}
