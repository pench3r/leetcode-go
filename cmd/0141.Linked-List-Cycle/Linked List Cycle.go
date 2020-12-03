package leetcode

import "github.com/pench3r/leetcode-go/structure"

func hasCycle(head *structure.ListNode) bool {
	fast := head
	slow := head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
