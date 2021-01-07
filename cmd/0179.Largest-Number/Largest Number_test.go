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
	res string
}

// problem: 环链表的实现
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{3, 6, 9, 1}},
			ans{"9631"},
		},
		{
			paras{[]int{1440, 7548, 4240, 6616, 733, 4712, 883, 8, 9576}},
			ans{"9576888375487336616471242401440"},
		},
		{
			paras{[]int{1}},
			ans{"1"},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, largestNumber(p.nums))
	}
}
