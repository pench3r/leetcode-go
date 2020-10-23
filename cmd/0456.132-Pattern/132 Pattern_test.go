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
	bl bool
}

func Test_TwoSum(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}},
			ans{false},
		},
		{
			paras{[]int{1, 2, 3, 4}},
			ans{false},
		},
		{
			paras{[]int{3, 1, 4, 2}},
			ans{true},
		},
		{
			paras{[]int{-1, 3, 2, 0}},
			ans{true},
		},
		{
			paras{[]int{3, 5, 0, 3, 4}},
			ans{true},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, find132pattern(p.nums))
	}

}