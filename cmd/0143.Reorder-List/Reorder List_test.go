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
	res []int
}

// problem: 环链表的实现
func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 2, 3, 4, 5}},
			ans{[]int{1, 5, 2, 4, 3}},
		},
		{
			paras{[]int{1, 2, 3, 4}},
			ans{[]int{1, 4, 2, 3}},
		},
		{
			paras{[]int{1}},
			ans{[]int{1}},
		},
		{
			paras{[]int{}},
			ans{[]int{}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		l1 := structure.Int2List(p.l1)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(reorderList(l1)))
	}
}
