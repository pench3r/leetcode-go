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
	res int
}

// problem: 环链表的实现
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{3, 6, 9, 1}},
			ans{3},
		},
		{
			paras{[]int{2, 435, 214, 64321, 643, 7234, 7, 436523, 7856, 8}},
			ans{372202},
		},
		{
			paras{[]int{1}},
			ans{0},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, maximumGap(p.nums))
	}
}
