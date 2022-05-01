package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := ListNode{Val: 1}
	curr1 := &list1
	for _, i := range []int{3} {
		curr1.Next = &ListNode{Val: i}
		curr1 = curr1.Next
	}
	list2 := ListNode{Val: 2}
	curr2 := &list2
	for _, i := range []int{4, 6} {
		curr2.Next = &ListNode{Val: i}
		curr2 = curr2.Next
	}
	res := mergeTwoLists(&list1, &list2)
	assert.Equal(t, res.toArray(), []int{1, 2, 3, 4, 6})
}
