package leetcode

import (
	"fmt"
	"testing"
)

type sample struct {
	paras
	ans
}

type paras struct {
	nums []int
}

type ans struct {
	res []int
}

// problem: 环链表的实现
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}},
			ans{[]int{}},
		},
		{
			paras{[]int{1, 5, 1, 1, 6, 4}},
			ans{[]int{1, 6, 1, 5, 1, 4}},
		},
		{
			paras{[]int{1, 3, 2, 2, 3, 1}},
			ans{[]int{2, 3, 1, 3, 1, 2}},
		},
		{
			paras{[]int{1, 1, 2, 1, 2, 2, 1}},
			ans{[]int{1, 2, 1, 2, 1, 2, 1}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		wiggleSort(p.nums)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, p.nums)
	}
}
