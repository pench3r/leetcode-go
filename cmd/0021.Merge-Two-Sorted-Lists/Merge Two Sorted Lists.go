package leetcode

import "github.com/pench3r/leetcode-go/structure"

func mergeTwoLists(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}
