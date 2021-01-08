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
	one string
}

type ans struct {
	res string
}

// problem: 环链表的实现
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{"aab"},
			ans{"aba"},
		},
		{
			paras{"aaab"},
			ans{""},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, reorganizeString(p.one))
	}
}
