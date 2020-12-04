package leetcode

import "github.com/pench3r/leetcode-go/structure"

func detectCycle(head *structure.ListNode) *structure.ListNode {
	fast := head
	slow := getCyclePoint(head)
	for slow != nil && fast != nil {
		if slow == fast {
			return slow
		}
		fast = fast.Next
		slow = slow.Next
	}
	return &structure.ListNode{Val: -1}
}

func getCyclePoint(head *structure.ListNode) *structure.ListNode {
	fast := head
	slow := head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return fast
		}
	}
	return nil
}
