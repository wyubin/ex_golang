package easyother

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head // find middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var p *ListNode // reverse from middle
	for slow != nil {
		slow, slow.Next, p = slow.Next, p, slow
	}
	for p != nil { // Compare
		if head.Val != p.Val {
			return false
		}
		head, p = head.Next, p.Next
	}
	return true
}
