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
	num int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{}},
			ans{0},
		},
		{
			paras{[]int{1, 2, 3}},
			ans{25},
		},
		{
			paras{[]int{4, 9, 0, 5, 1}},
			ans{1026},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, sumNumbers(root))
	}

}
