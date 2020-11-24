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
	target int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 3, 4, 2, 2}},
			ans{2},
		},
		{
			paras{[]int{3, 1, 3, 4, 2}},
			ans{3},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, findDuplicate(p.nums))
	}

}
