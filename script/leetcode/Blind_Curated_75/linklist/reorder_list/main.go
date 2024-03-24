package main

import "fmt"

func reorder(list *ListNode) *ListNode {
	// split list
	fastNode, slowNode := list, list
	isSlowRun := false
	for fastNode.Next != nil {
		fastNode = fastNode.Next
		if isSlowRun {
			slowNode = slowNode.Next
		}
		isSlowRun = !isSlowRun
	}

	// reverse secode list
	var prev *ListNode
	curr := slowNode.Next
	for curr != nil {
		tmpNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNode
	}
	// close first list
	slowNode.Next = nil

	// insert
	head1, head2 := list, prev
	for head2 != nil {
		fmt.Printf("head1[%d], head2[%d]\n", head1.Val, head2.Val)
		tmpNode := head1.Next
		head1.Next = head2
		head1 = head2
		head2 = tmpNode
	}

	return list
}

func main() {
	s := []int{1, 2, 3, 4}
	k := reorder(slice2ListNode(s, -1))
	fmt.Printf("k: %s\n", k)
}
