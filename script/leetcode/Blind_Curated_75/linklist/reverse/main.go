package main

import "fmt"

func reverse(list *ListNode) *ListNode {
	var prev *ListNode
	var tmpNode *ListNode
	curr := list
	for curr != nil {
		tmpNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNode
	}
	return prev
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	k := reverse(slice2ListNode(s, -1))
	fmt.Printf("k: %s\n", k)
}
