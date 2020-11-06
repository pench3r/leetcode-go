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
	num []int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 2, 5, 3, 4, structure.NULL, 6}},
			ans{[]int{1, structure.NULL, 2, structure.NULL, 3, structure.NULL, 4, structure.NULL, 5, structure.NULL, 6}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		flatten(root)
		fmt.Printf("[except]: %v, [output]: %v\n", a, structure.Tree2Preorder(root))
	}

}
