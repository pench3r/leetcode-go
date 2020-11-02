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
			paras{[]int{10, 5, 15, structure.NULL, structure.NULL, 6, 20}},
			ans{false},
		},
		{
			paras{[]int{}},
			ans{true},
		},
		{
			paras{[]int{2, 1, 3}},
			ans{true},
		},
		{
			paras{[]int{5, 1, 4, structure.NULL, structure.NULL, 3, 6}},
			ans{false},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, isValidBST1(root))
	}

}
