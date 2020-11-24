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
	target int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{"12"},
			ans{2},
		},
		{
			paras{"226"},
			ans{3},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v,  [output]: %v\n", a, numDecodings(p.s))
	}

}
