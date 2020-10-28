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
			paras{[]int{0,1,0,2,1,0,1,3,2,1,2,1}},
			ans{6},
		},
		{
			paras{[]int{4,2,0,3,2,5}},
			ans{9},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, trap(p.nums))
	}

}