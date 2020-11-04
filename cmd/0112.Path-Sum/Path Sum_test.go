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
	two int
}

type ans struct {
	bl bool
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}, 0},
			ans{false},
		},
		{
			paras{[]int{5, 4, 8, 11, structure.NULL, 13, 4, 7, 2, structure.NULL, structure.NULL, structure.NULL, 1}, 22},
			ans{true},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, hasPathSum(root, p.two))
	}

}
