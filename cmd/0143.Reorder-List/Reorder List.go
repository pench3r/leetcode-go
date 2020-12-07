package leetcode

import "github.com/pench3r/leetcode-go/structure"

func reorderList(head *structure.ListNode) *structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	preMidHead := slow
	preCurNode := slow.Next
	for preCurNode.Next != nil {
		curNode := preCurNode.Next
		preCurNode.Next = curNode.Next
		curNode.Next = preMidHead.Next
		preMidHead.Next = curNode
	}
	p1 := head
	p2 := preMidHead.Next
	for p1 != preMidHead {
		preMidHead.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = preMidHead.Next
	}
	return head
}
