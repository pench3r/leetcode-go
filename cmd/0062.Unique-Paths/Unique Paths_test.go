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
	m int
	n int
}

type ans struct {
	target int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{3, 2},
			ans{3},
		},
		{
			paras{7, 3},
			ans{28},
		},
		{
			paras{1, 2},
			ans{1},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, uniquePaths(p.m, p.n))
	}

}
