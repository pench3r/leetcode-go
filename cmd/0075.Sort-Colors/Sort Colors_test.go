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
	target []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{2, 0, 2, 1, 1, 0}},
			ans{[]int{0, 0, 1, 1, 2, 2}},
		},
		{
			paras{[]int{2, 0, 1, 1, 2, 0, 2, 1, 2, 0, 0, 0, 1, 2, 2, 2, 0, 1, 1}},
			ans{[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2}},
		},
		{
			paras{[]int{1}},
			ans{[]int{1}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		sortColors(p.nums)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, p.nums)
	}

}
