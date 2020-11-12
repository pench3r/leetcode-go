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
	bw string
	ew string
	wl []string
}

type ans struct {
	num int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			ans{5},
		},
		{
			paras{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			ans{0},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		fmt.Printf("[except]: %v, [output]: %v\n", a, ladderLength(p.bw, p.ew, p.wl))
	}

}
