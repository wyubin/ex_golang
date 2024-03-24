package main

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
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
	return head
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	n := 2
	head := slice2ListNode(nums, -1)
	fmt.Printf("head:%s\n", head)
	k := removeNthFromEnd(head, n)
	fmt.Printf("k:%s\n", k)
}
