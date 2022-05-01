package easyother

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) toArray() []int {
	res := []int{}
	curr := s
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	return res
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 先做一個空的頭
	dummy := ListNode{}
	curr := &dummy
	// 在兩個 list 都還在 node 上時，可以用比較來修改next
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	// 如果其中已經到 nil，那剩下的就直接接到後面
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	fmt.Printf("array:%+v\n", dummy.Next.toArray())
	return dummy.Next
}
