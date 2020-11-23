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
	og [][]int
}

type ans struct {
	target int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			ans{2},
		},
		{
			paras{[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			}},
			ans{0},
		},
		{
			paras{[][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			}},
			ans{0},
		},
		{
			paras{[][]int{
				{1, 0},
			}},
			ans{0},
		},
		{
			paras{[][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			}},
			ans{0},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, uniquePathsWithObstacles(p.og))
	}

}
