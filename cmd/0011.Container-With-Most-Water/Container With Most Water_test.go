package leetcode

import (
	"testing";
	"fmt";
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
			paras{[]int{1, 1}},
			ans{1},
		},
		{
			paras{[]int{1,8,6,2,5,4,8,3,7}},
			ans{49},
		},
		{
			paras{[]int{4,3,2,1,4}},
			ans{16},
		},
		{
			paras{[]int{1,2,1}},
			ans{2},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, maxArea(p.nums))
	}

}