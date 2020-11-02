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
}

type ans struct {
	bl bool
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{3, 4, 4, 5, structure.NULL, structure.NULL, 5, 6, structure.NULL, structure.NULL, 6}},
			ans{true},
		},
		{
			paras{[]int{}},
			ans{true},
		},
		{
			paras{[]int{1, 2, 2, structure.NULL, 3, 3}},
			ans{true},
		},
		{
			paras{[]int{1}},
			ans{true},
		},
		{
			paras{[]int{1, 2, 3}},
			ans{false},
		},
		{
			paras{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans{true},
		},
		{
			paras{[]int{1, 2, 2, structure.NULL, 3, structure.NULL, 3}},
			ans{false},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, isSymmetric(root))
	}

}
