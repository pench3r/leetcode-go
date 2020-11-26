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
	ans int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 2, 3, 4, 5, 6}},
			ans{6},
		},
		{
			paras{[]int{1, 2, 3}},
			ans{3},
		},
		{
			paras{[]int{1}},
			ans{1},
		},
		{
			paras{[]int{}},
			ans{0},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, countNodes(root))
	}

}
