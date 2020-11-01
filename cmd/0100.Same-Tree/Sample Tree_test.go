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
	one []int
	two []int
}

type ans struct {
	bl bool
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}, []int{}},
			ans{true},
		},
		{
			paras{[]int{}, []int{1}},
			ans{false},
		},
		{
			paras{[]int{1}, []int{1}},
			ans{true},
		},
		{
			paras{[]int{1, 2, 3}, []int{1, 2, 3}},
			ans{true},
		},
		{
			paras{[]int{1, 2}, []int{1, structure.NULL, 2}},
			ans{false},
		},
		{
			paras{[]int{1, 2, 1}, []int{1, 1, 2}},
			ans{false},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		pt, qt := structure.Ints2TreeNode(p.one), structure.Ints2TreeNode(p.two)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, isSameTree(pt, qt))
	}

}
