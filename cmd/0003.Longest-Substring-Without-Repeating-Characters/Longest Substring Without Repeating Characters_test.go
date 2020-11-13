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
}

type ans struct {
	res int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{"abcabcbb"},
			ans{3},
		},
		{
			paras{"bbbbb"},
			ans{1},
		},
		{
			paras{"pwwkew"},
			ans{3},
		},
		{
			paras{""},
			ans{0},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, lengthOfLongestSubstring(p.s))
	}

}
