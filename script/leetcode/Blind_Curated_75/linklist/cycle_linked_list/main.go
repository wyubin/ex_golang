package main

import "fmt"

func isCycle(list *ListNode) bool {
	fastNode := list
	slowNode := list
	isSlowRun := false
	for fastNode != nil {
		fastNode = fastNode.Next
		if isSlowRun {
			slowNode = slowNode.Next
		}
		isSlowRun = !isSlowRun // swap
		if fastNode == slowNode {
			return true
		}
	}
	return false
}

func main() {
	s := []int{3, 2, 0, -4}
	list := slice2ListNode(s, 1)
	fmt.Printf("list: %s\n", list)
	k := isCycle(list)
	fmt.Printf("k: %+v\n", k)
}
