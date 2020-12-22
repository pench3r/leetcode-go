package leetcode

import (
	"github.com/pench3r/leetcode-go/structure"
)

func getIntersectionNode(headA, headB *structure.ListNode) *structure.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
