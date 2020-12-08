package leetcode

import "github.com/pench3r/leetcode-go/structure"

func reverseList(head *structure.ListNode) *structure.ListNode {
	var behind *structure.ListNode

	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}
