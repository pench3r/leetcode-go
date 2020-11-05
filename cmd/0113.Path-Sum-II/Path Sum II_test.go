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
	num [][]int
}

func Test_Problem(t *testing.T) {
	ss := []sample{
		{
			paras{[]int{1, 0, 1, 1, 2, 0, -1, 0, 1, -1, 0, -1, 0, 1, 0}, 2},
			ans{[][]int{{1, 0, 1, 0}, {1, 0, 2, -1}, {1, 1, 0, 0}, {1, 1, -1, 1}}},
		},
		{
			paras{[]int{}, 0},
			ans{[][]int{{}}},
		},
		{
			paras{[]int{5, 4, 8, 11, structure.NULL, 13, 4, 7, 2, structure.NULL, structure.NULL, 5, 1}, 22},
			ans{[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		},
	}

	for _, s := range ss {
		p, a := s.paras, s.ans
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("[except]: %v,  [output]: %v\n", a, pathSum(root, p.two))
	}

}
