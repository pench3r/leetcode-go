package leetcode

import "github.com/pench3r/leetcode-go/structure"

func partition(head *structure.ListNode, x int) *structure.ListNode {
	beforeHead := &structure.ListNode{}
	before := beforeHead
	afterHead := &structure.ListNode{}
	after := afterHead
	for head != nil {
		if head.Val < x {
			beforeHead.Next = head
			beforeHead = beforeHead.Next
		} else {
			afterHead.Next = head
			afterHead = afterHead.Next
		}
		head = head.Next
	}
	afterHead.Next = nil
	beforeHead.Next = after.Next
	return before.Next
}
