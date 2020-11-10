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
	one    int
	target int
}

type ans struct {
	res [][]int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{3, 7},
			ans{[][]int{{1, 2, 4}}},
		},
		{
			paras{3, 9},
			ans{[][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, combinationSum3(p.one, p.target))
	}

}
