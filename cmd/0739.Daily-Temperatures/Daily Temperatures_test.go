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
	num []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
			ans{[]int{1, 1, 4, 2, 1, 1, 0, 0}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, dailyTemperatures(p.nums))
	}

}