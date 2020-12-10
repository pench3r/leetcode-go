package leetcode

import "github.com/pench3r/leetcode-go/structure"

func reverseKGroup(head *structure.ListNode, k int) *structure.ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *structure.ListNode, last *structure.ListNode) *structure.ListNode {
	prev := last
	for first != last {
		next := first.Next
		first.Next = prev
		prev = first
		first = next
	}
	return prev
}
