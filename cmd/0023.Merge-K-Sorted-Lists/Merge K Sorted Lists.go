package leetcode

import "github.com/pench3r/leetcode-go/structure"

func mergeKLists(lists []*structure.ListNode) *structure.ListNode {
	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*structure.ListNode, start int, end int) *structure.ListNode {
	if start <= end {
		if start == end {
			return lists[start]
		}
		if end-start == 1 {
			return mergeTwoLists(lists[start], lists[end])
		}
		mid := start + (end-start+1)/2
		leftList := mergeLists(lists, start, mid)
		rightList := mergeLists(lists, mid+1, end)
		return mergeTwoLists(leftList, rightList)
	}
	return nil
}

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
