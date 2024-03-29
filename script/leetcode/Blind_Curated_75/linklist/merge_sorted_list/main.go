package main

import "fmt"

func solution(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := ListNode{}
	curr := &dummy

	for list1 != nil && list2 != nil {
		// 比較兩個 list head 的值，如果比較小就接上去，並往下一個移動
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}
	return dummy.Next
}

func main() {
	list1 := slice2ListNode([]int{1, 2, 4}, -1)
	list2 := slice2ListNode([]int{1, 3, 4}, -1)
	// fmt.Printf("k: %s\n", solution(list1, list2))
	// merge k liked list
	all_list := []*ListNode{list2, slice2ListNode([]int{1, 2, 5}, -1)}
	res := &ListNode{Next: list1}
	for i := 0; i < len(all_list); i++ {
		res = solution(res, all_list[i])
	}
	fmt.Printf("k: %s\n", res.Next)
}
