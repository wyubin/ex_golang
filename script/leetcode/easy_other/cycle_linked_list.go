package easyother

func hasCycle(head *ListNode) bool {
	// two pointer
	slow := head
	fast := head
	stepSlowRun := false

	for fast != nil {
		fast = fast.Next
		if stepSlowRun {
			slow = slow.Next
		}
		// taggle step slow
		stepSlowRun = !stepSlowRun
		// reach same
		if fast != nil && fast == slow {
			return true
		}
	}

	return false
}
