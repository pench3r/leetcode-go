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
	num	[][]int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1}},
			ans{[][]int{{1}}},
		},
		{
			paras{[]int{}},
			ans{[][]int{}},
		},
		{
			paras{[]int{3, 9, 20, structure.NULL, structure.NULL, 15, 7}},
			ans{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, levelOrder(root))
	}

}