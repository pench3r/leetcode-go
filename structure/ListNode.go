package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func Int2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNode{}
	t := l
	for _, num := range nums {
		t.Next = &ListNode{Val: num}
		t = t.Next
	}
	return l.Next
}
