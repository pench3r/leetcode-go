package structure

import "fmt"

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

func List2Int(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	limit, time := 100, 0
	cur, res := head, []int{}
	for cur != nil {
		if time > limit {
			msg := fmt.Sprintf("链条深度超过%d, 可能出现环链", limit)
			panic(msg)
		}
		res = append(res, cur.Val)
		cur = cur.Next
		time++
	}
	return res
}
