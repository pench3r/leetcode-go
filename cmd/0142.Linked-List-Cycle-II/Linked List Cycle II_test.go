package leetcode

import (
	"fmt"
	"testing"

	"github.com/pench3r/leetcode-go/structure"
)

type sample struct {
	paras
	ans
}

type paras struct {
	l1 []int
}

type ans struct {
	res int
}

// problem: 环链表的实现
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{3, 2, 0, -4}},
			ans{1},
		},
		{
			paras{[]int{1, 2}},
			ans{0},
		},
		{
			paras{[]int{1}},
			ans{-1},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		l1 := structure.Int2List(p.l1)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, detectCycle(l1).Val)
	}

}
