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
	one    []int
	target int
}

type ans struct {
	res [][]int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{2, 3, 6, 7}, 7},
			ans{[][]int{{7}, {2, 2, 3}}},
		},
		{
			paras{[]int{2, 3, 5}, 8},
			ans{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, combinationSum(p.one, p.target))
	}

}
