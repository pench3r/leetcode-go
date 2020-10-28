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
	nums [][]int
}

type ans struct {
	num int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
			ans{2},
		},
		{
			paras{[][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}},
			ans{1},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, findCircleNum(p.nums))
	}

}
