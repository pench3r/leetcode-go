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
	target int
}

type ans struct {
	nums []int
}

func Test_TwoSum(t *testing.T) {
	ss := []sample{
		{
			paras{
				[]int{3,2,4},
				6,
			},
			ans{
				[]int{1,2},
			},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, twoSum(p.nums, p.target))
	}

}