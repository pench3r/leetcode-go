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
	num int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{
				[]int{2, 1, 5, 6, 2, 3},
			},
			ans{
				10,
			},
		},
		{
			paras{
				[]int{1},
			},
			ans{
				1,
			},
		},
		{
			paras{
				[]int{1,1},
			},
			ans{
				2,
			},
		},
		{
			// bug1
			paras{
				[]int{4, 2, 0, 3, 2, 5},
			},
			ans{
				6,
			},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, largestRectangleArea(p.nums))
	}

}