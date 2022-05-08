package easyother

func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil {
		head, head.Next, p = head.Next, p, head
		// head 會換到下一個，並且 head.next 會接到 p，p再接目前的 head
	}
	return p
}
