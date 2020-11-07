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
	ans []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 3, structure.NULL, structure.NULL, 2}},
			ans{[]int{3, 1, structure.NULL, structure.NULL, 2}},
		},
		{
			paras{[]int{3, 1, 4, structure.NULL, structure.NULL, 2}},
			ans{[]int{2, 1, 4, structure.NULL, structure.NULL, 3}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		recoverTree(root)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, structure.Tree2Preorder(root))
	}

}
