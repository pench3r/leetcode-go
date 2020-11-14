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
	s string
	p string
}

type ans struct {
	res string
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{"ADOBECODEBANC", "ABC"},
			ans{"BANC"},
		},
		{
			paras{"a", "aa"},
			ans{""},
		},
		{
			paras{"a", "a"},
			ans{"a"},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, minWindow(p.s, p.p))
	}

}
