package leetcode

import "github.com/pench3r/leetcode-go/structure"

func sortList(head *structure.ListNode) *structure.ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}
	midNode := middleNode(head)
	cur = midNode.Next
	midNode.Next = nil
	midNode = cur
	left := sortList(head)
	right := sortList(midNode)
	return mergeTwoLists(left, right)
}

func middleNode(head *structure.ListNode) *structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head
	slow := head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoLists(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}
