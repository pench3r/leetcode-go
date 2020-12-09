package leetcode

import "github.com/pench3r/leetcode-go/structure"

func insertionSortList(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return nil
	}
	sorted := &structure.ListNode{Val: 0, Next: nil}
	cur, prev := head, sorted
	for cur != nil {
		next := cur.Next
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		cur.Next = prev.Next
		prev.Next = cur
		prev = sorted
		cur = next
	}
	return sorted.Next
}
