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
	l2 []int
}

type ans struct {
	target []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}, []int{}},
			ans{[]int{}},
		},
		{
			paras{[]int{1}, []int{1}},
			ans{[]int{1, 1}},
		},
		{
			paras{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},
		{
			paras{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},
		{
			paras{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans{[]int{1, 9, 9, 9, 9, 9}},
		},
		{
			paras{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans{[]int{1, 9, 9, 9, 9, 9}},
		},
		{
			paras{[]int{2, 3, 4}, []int{4, 5, 6}},
			ans{[]int{2, 3, 4, 4, 5, 6}},
		},
		{
			paras{[]int{1, 3, 8}, []int{1, 7}},
			ans{[]int{1, 1, 3, 7, 8}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		l1 := structure.Int2List(p.l1)
		l2 := structure.Int2List(p.l2)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.List2Int(mergeTwoLists(l1, l2)))
	}

}
