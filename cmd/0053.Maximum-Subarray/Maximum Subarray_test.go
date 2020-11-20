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
			paras{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans{6},
		},
		{
			paras{[]int{2, 7, 9, 3, 1}},
			ans{22},
		},
		{
			paras{[]int{-1, -2}},
			ans{-1},
		},
		{
			paras{[]int{2}},
			ans{2},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, maxSubArray(p.nums))
	}

}
