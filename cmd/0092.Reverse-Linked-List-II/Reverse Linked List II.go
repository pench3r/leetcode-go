package leetcode

import "github.com/pench3r/leetcode-go/structure"

func reverseBetween(head *structure.ListNode, m int, n int) *structure.ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &structure.ListNode{Val: 0, Next: head}
	prev := newHead
	for count := 0; prev.Next != nil && count < m-1; count++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return head
	}
	cur := prev.Next
	for i := 0; i < n-m; i++ {
		next := prev.Next
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = next
	}
	return newHead.Next
}
